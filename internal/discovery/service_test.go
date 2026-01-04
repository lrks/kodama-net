package discovery_test

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/lrks/kodama-net/internal/discovery"
	"github.com/lrks/kodama-net/internal/echonetlite"
)

// fakeConn は discovery.Conn を満たすテスト用の実装です。
type fakeConn struct {
	writtenPayloads [][]byte
	writtenAddrs    []*net.UDPAddr

	// Read 側で返すデータとエラー
	reads []fakeRead
	idx   int
}

type fakeRead struct {
	data []byte
	addr *net.UDPAddr
	err  error
}

func (f *fakeConn) ReadFromUDP(b []byte) (int, *net.UDPAddr, error) {
	if f.idx >= len(f.reads) {
		// これ以上読むものがない場合はタイムアウト相当のエラーを返す
		return 0, nil, &net.DNSError{IsTimeout: true}
	}
	r := f.reads[f.idx]
	f.idx++
	copy(b, r.data)
	return len(r.data), r.addr, r.err
}

func (f *fakeConn) WriteToUDP(b []byte, addr *net.UDPAddr) (int, error) {
	f.writtenPayloads = append(f.writtenPayloads, append([]byte(nil), b...))
	f.writtenAddrs = append(f.writtenAddrs, addr)
	return len(b), nil
}

func (f *fakeConn) SetReadDeadline(t time.Time) error {
	_ = t
	return nil
}

// dummyRequestPayload が妥当な ECHONET Lite フレームとしてパースできることだけ
// 確認しておく。実際のネットワーク I/O はここではテストしない。
func TestDummyRequestPayload_IsValidECHONETLiteFrame(t *testing.T) {
	// dummyRequestPayload は discovery パッケージ内部の未公開変数なので、
	// ここではサービス経由でアクセスせず、同じバイト列を生成して検証する。
	payload := []byte{
		0x10, 0x81,
		0x00, 0x01,
		0x05, 0xff, 0x01,
		0x0e, 0xf0, 0x01,
		0x62,
		0x01,
		0xd6,
		0x00,
	}

	if _, err := echonetlite.Parse(payload); err != nil {
		t.Fatalf("dummy discovery payload must be valid ECHONET Lite frame, but got error: %v", err)
	}

	_ = discovery.NewService() // 型チェック用に一度だけ呼んでおく
}

func TestDiscover_SendsMulticastRequestAndCollectsIPs(t *testing.T) {
	svc := discovery.NewService()
	fc := &fakeConn{
		reads: []fakeRead{{
			data: buildDummyResponse(),
			addr: &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 3610},
		}, {
			// タイムアウトを示す
			data: nil,
			addr: nil,
			// net.Error で Timeout()==true を返すものを使う
			// ここでは DNS エラーを使うが、テスト目的では十分。
			// discovery 側では net.Error として扱われる。
			err: &net.DNSError{IsTimeout: true},
		}},
	}

	ctx := context.Background()
	ipTimeout := 1 * time.Second

	ips, err := svc.Discover(ctx, fc, ipTimeout)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Discover returned unexpected error: %v", err)
	}

	if len(fc.writtenPayloads) != 1 {
		t.Fatalf("expected 1 written packet, got %d", len(fc.writtenPayloads))
	}

	if len(fc.writtenAddrs) != 1 {
		t.Fatalf("expected 1 written addr, got %d", len(fc.writtenAddrs))
	}

	addr := fc.writtenAddrs[0]
	if addr.Port != 3610 || !addr.IP.Equal(net.IPv4(224, 0, 23, 0)) {
		t.Fatalf("unexpected multicast addr: %+v", addr)
	}

	if _, err := echonetlite.Parse(fc.writtenPayloads[0]); err != nil {
		t.Fatalf("written payload must be valid ECHONET Lite frame, got error: %v", err)
	}

	if len(ips) != 1 || !ips[0].Equal(net.IPv4(192, 0, 2, 1)) {
		t.Fatalf("unexpected discovered IPs: %#v", ips)
	}
}

func TestProbe_SendsRequestsToEachIP(t *testing.T) {
	svc := discovery.NewService()
	fc := &fakeConn{
		reads: []fakeRead{
			{data: buildDummyResponse(), addr: &net.UDPAddr{IP: net.IPv4(192, 0, 2, 1), Port: 3610}},
			{data: buildDummyResponse(), addr: &net.UDPAddr{IP: net.IPv4(192, 0, 2, 2), Port: 3610}},
		},
	}

	ctx := context.Background()
	timeout := 1 * time.Second
	ips := []net.IP{net.IPv4(192, 0, 2, 1), net.IPv4(192, 0, 2, 2)}

	if err := svc.Probe(ctx, fc, ips, timeout); err != nil {
		t.Fatalf("Probe returned error: %v", err)
	}

	if len(fc.writtenAddrs) != len(ips) {
		t.Fatalf("expected %d written addrs, got %d", len(ips), len(fc.writtenAddrs))
	}

	for i, ip := range ips {
		addr := fc.writtenAddrs[i]
		if addr.Port != 3610 || !addr.IP.Equal(ip) {
			t.Fatalf("written addr[%d] mismatch: got %+v, want IP=%v, Port=3610", i, addr, ip)
		}
	}
}

// buildDummyResponse は簡易な ECHONET Lite レスポンス電文を構築する。
// Discover/Probe の中ではパース結果は使っていないため、最低限パース可能
// であれば問題ない。
func buildDummyResponse() []byte {
	return []byte{
		0x10, 0x81,
		0x12, 0x34,
		0x05, 0xff, 0x01,
		0x0e, 0xf0, 0x01,
		0x72,
		0x01,
		0xd6,
		0x01,
		0x00,
	}
}
