package discovery

import (
	"context"
	"errors"
	"fmt"
	"net"
	"slices"
	"time"

	"github.com/lrks/kodama-net/internal/echonetlite"
)

type Device struct {
	IPAddr net.IP
	EOJ    [3]byte
}

type Conn interface {
	ReadFromUDP(b []byte) (int, *net.UDPAddr, error)
	WriteToUDP(b []byte, addr *net.UDPAddr) (int, error)
	SetReadDeadline(t time.Time) error
}

type Service interface {
	// 探索を実行し、応答したデバイス一覧を返す
	Discover(ctx context.Context, conn Conn, timeout time.Duration) ([]Device, error)

	// 指定されたデバイスの各種プロパティを取得する
	Probe(ctx context.Context, conn Conn, device Device, timeout time.Duration) ([]echonetlite.Property, error)

	// 指定されたデバイスのクラス定義を取得する
	GetClassDefinition(device Device) (echonetlite.ClassDefinition, error)

	// 指定されたデバイスのプロパティ定義を取得する
	GetPropertyDefinitionFromMap(device Device, properties []echonetlite.Property, targetEPC byte) ([]echonetlite.PropertyDefinition, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

const maxPayloadSize = 1472

// 探索.
func (s *service) Discover(ctx context.Context, conn Conn, timeout time.Duration) ([]Device, error) {
	// リクエスト送信
	var requestPayload = []byte{
		0x10, 0x81, // EHD1, EHD2
		0x12, 0x34, // TID
		0x05, 0xff, 0x01, // SEOJ (Controller)
		0x0e, 0xf0, 0x01, // DEOJ (NodeProfile)
		0x62, // ESV (Get)
		0x01, // OPC
		0xd6, // EPC (自ノードインスタンスリストS)
		0x00, // PDC
	}

	maddr, err := net.ResolveUDPAddr("udp4", "224.0.23.0:3610")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve multicast address: %w", err)
	}

	if _, err := conn.WriteToUDP(requestPayload, maddr); err != nil {
		return nil, fmt.Errorf("failed to send discovery request: %w", err)
	}

	// レスポンス受信
	var (
		buf     = make([]byte, maxPayloadSize)
		devices []Device
	)

	deadline := time.Now().Add(timeout)
	if err := conn.SetReadDeadline(deadline); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return devices, ctx.Err() //nolint:wrapcheck
		default:
		}

		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			var nerr net.Error
			if errors.As(err, &nerr) && nerr.Timeout() {
				// Timeout
				return devices, nil
			}

			return devices, fmt.Errorf("failed to read from UDP: %w", err)
		}

		// 無効なレスポンスは無視
		if n == 0 || addr == nil || addr.IP == nil || addr.IP.IsLoopback() || addr.IP.IsMulticast() || addr.IP.IsUnspecified() {
			continue
		}

		frame, err := echonetlite.ParseFrame(buf[:n])
		if err != nil {
			// パース失敗
			continue
		}

		if frame.ESV != 0x72 {
			// 応答がGet_Resではない
			continue
		}

		if len(frame.Properties) != 1 {
			// プロパティが1つでない
			continue
		}

		if frame.Properties[0].EPC != 0xd6 {
			// プロパティが想定外
			continue
		}

		if frame.Properties[0].PDC == 0 {
			// データなし
			continue
		}

		// 追加していく
		devices = append(devices, Device{IPAddr: addr.IP, EOJ: frame.SEOJ})

		count := int(frame.Properties[0].EDT[0])
		for i := range count {
			offset := 1 + i*3
			if offset+3 > len(frame.Properties[0].EDT) {
				break
			}

			devices = append(devices, Device{IPAddr: addr.IP, EOJ: [3]byte{
				frame.Properties[0].EDT[offset],
				frame.Properties[0].EDT[offset+1],
				frame.Properties[0].EDT[offset+2],
			}})
		}
	}

	return devices, nil
}

// 詳細取得.
func (s *service) Probe(ctx context.Context, conn Conn, device Device, timeout time.Duration) ([]echonetlite.Property, error) {
	var properties []echonetlite.Property

	// Getプロパティマップの取得
	getPropertyMap, err := s.get(ctx, conn, device, echonetlite.GetPropertyMapEPC, timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to get property map from %s,%02x%02x%02x: %w", device.IPAddr, device.EOJ[0], device.EOJ[1], device.EOJ[2], err)
	}

	properties = append(properties, getPropertyMap)

	// GetプロパティマップからEPCを抽出
	epcs, err := echonetlite.ParsePropertyMap(getPropertyMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse property map: %w", err)
	}

	// EPCをGetしていく
	for _, epc := range epcs {
		if epc == echonetlite.GetPropertyMapEPC {
			// Getプロパティマップは取得済み
			continue
		}

		prop, err := s.get(ctx, conn, device, epc, timeout)
		if err != nil {
			return nil, fmt.Errorf("failed to get property %02x from %s,%02x%02x%02x: %w", epc, device.IPAddr, device.EOJ[0], device.EOJ[1], device.EOJ[2], err)
		}

		properties = append(properties, prop)
	}

	// ソートする
	slices.SortFunc(properties, func(i, j echonetlite.Property) int {
		return int(i.EPC) - int(j.EPC)
	})

	return properties, nil
}

// デバイスのクラス定義を取得する.
func (s *service) GetClassDefinition(device Device) (echonetlite.ClassDefinition, error) {
	key := [2]byte{device.EOJ[0], device.EOJ[1]}
	if _, ok := echonetlite.ClassDefinitions[key]; !ok {
		return echonetlite.ClassDefinition{}, fmt.Errorf("class definition not found for EOJ %02x%02x%02x", device.EOJ[0], device.EOJ[1], device.EOJ[2])
	}

	classDefinition := echonetlite.ClassDefinitions[key]

	// スーパークラスも含めて返す
	superClassKey := [2]byte{0x00, 0x00}
	if _, ok := echonetlite.ClassDefinitions[superClassKey]; !ok {
		return echonetlite.ClassDefinition{}, fmt.Errorf("super class definition not found for EOJ %02x%02x", superClassKey[0], superClassKey[1])
	}

	superClassDefinition := echonetlite.ClassDefinitions[superClassKey]
	merged := make([]echonetlite.PropertyDefinition, 0, len(classDefinition.Properties)+len(superClassDefinition.Properties))
	merged = append(merged, classDefinition.Properties...)
	merged = append(merged, superClassDefinition.Properties...)
	classDefinition.Properties = merged

	return classDefinition, nil
}

// 指定されたEPCのプロパティマップからプロパティ定義を取得する.
func (s *service) GetPropertyDefinitionFromMap(device Device, properties []echonetlite.Property, targetEPC byte) ([]echonetlite.PropertyDefinition, error) {
	// クラス定義
	classDefinition, err := s.GetClassDefinition(device)
	if err != nil {
		return nil, fmt.Errorf("failed to get class definition: %w", err)
	}

	// 規格Versionの取得
	var appendixVersion byte = 0x00

	if !classDefinition.IsNodeProfile {
		for _, prop := range properties {
			if prop.EPC == echonetlite.VersionEPC && prop.PDC == 4 {
				appendixVersion = prop.EDT[2]

				break
			}
		}
	}

	// プロパティマップの取得
	foundPropertyMap := false

	var propertyMap echonetlite.Property

	for _, prop := range properties {
		if prop.EPC == targetEPC {
			propertyMap = prop
			foundPropertyMap = true

			break
		}
	}

	if !foundPropertyMap {
		return nil, fmt.Errorf("property map not found for EPC %02x", targetEPC)
	}

	// プロパティマップからEPCを抽出
	epcs, err := echonetlite.ParsePropertyMap(propertyMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse property map: %w", err)
	}
	slices.Sort(epcs)

	// プロパティ定義の取得
	var propertyDefinitions []echonetlite.PropertyDefinition
	for _, epc := range epcs {
		if epc>>4 == 0xf {
			propertyDefinitions = append(propertyDefinitions, echonetlite.PropertyDefinition{
				EPC:           epc,
				Name:          "ユーザ定義領域",
				NameEN:        "User defined area",
				ShortName:     "userDefinedArea",
				Description:   "ユーザ定義領域",
				DescriptionEN: "User defined area",
			})
			continue
		}

		found := false
		for _, prop := range classDefinition.Properties {
			if epc == prop.EPC {
				if appendixVersion == 0x00 || (prop.ValidRelease.FROM <= appendixVersion && appendixVersion <= prop.ValidRelease.TO) {
					propertyDefinitions = append(propertyDefinitions, prop)
					found = true
					break
				}
			}
		}

		if !found {
			propertyDefinitions = append(propertyDefinitions, echonetlite.PropertyDefinition{
				EPC:           epc,
				Name:          "未定義",
				NameEN:        "Undefined",
				ShortName:     "undefined",
				Description:   "定義なし",
				DescriptionEN: "Undefined",
			})
		}
	}

	return propertyDefinitions, nil
}

func (s *service) get(ctx context.Context, conn Conn, device Device, epc byte, timeout time.Duration) (echonetlite.Property, error) {
	// リクエスト送信
	var requestPayload = []byte{
		0x10, 0x81, // EHD1, EHD2
		epc, 0x00, // TID (区別のためにEPCを埋めておく)
		0x05, 0xff, 0x01, // SEOJ (Controller)
		device.EOJ[0], device.EOJ[1], device.EOJ[2], // DEOJ
		0x62, // ESV (Get)
		0x01, // OPC
		epc,  // EPC
		0x00, // PDC
	}

	addr := &net.UDPAddr{IP: device.IPAddr, Port: 3610}

	// TODO: TTL=1で送信したほうが安全
	// 万が一device.IPAddrがLAN外のアドレスに偽装されても届かず、被害がLAN内で収まる
	if _, err := conn.WriteToUDP(requestPayload, addr); err != nil {
		return echonetlite.Property{}, fmt.Errorf("failed to send GET request: %w", err)
	}

	// レスポンス受信
	var buf = make([]byte, maxPayloadSize)

	deadline := time.Now().Add(timeout)
	if err := conn.SetReadDeadline(deadline); err != nil {
		return echonetlite.Property{}, fmt.Errorf("failed to set read deadline: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return echonetlite.Property{}, ctx.Err() //nolint:wrapcheck
		default:
		}

		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			return echonetlite.Property{}, fmt.Errorf("failed to read from UDP: %w", err)
		}

		// 無効なレスポンスは無視
		if n == 0 || addr == nil || addr.IP == nil || !addr.IP.Equal(device.IPAddr) {
			continue
		}

		// レスポンスをパース
		frame, err := echonetlite.ParseFrame(buf[:n])
		if err != nil {
			// パース失敗
			continue
		}

		if frame.ESV != 0x72 {
			// 応答がGet_Resではない
			continue
		}

		if len(frame.Properties) != 1 {
			// プロパティが1つでない
			continue
		}

		if frame.Properties[0].EPC != epc {
			// プロパティが想定外
			continue
		}

		// OK
		return frame.Properties[0], nil
	}

	// 応答が来ない場合もエラーにはしない
}
