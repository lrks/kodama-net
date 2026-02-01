package echonetlite

import (
	"encoding/binary"
	"errors"
)

func ParseFrame(b []byte) (*Frame, error) {
	if len(b) < 12 {
		return nil, errors.New("frame too short in header")
	}

	if b[0] != 0x10 || b[1] != 0x81 {
		return nil, errors.New("unsupported header")
	}

	f := &Frame{
		EHD1: b[0],
		EHD2: b[1],
		TID:  binary.BigEndian.Uint16(b[2:4]),
		SEOJ: [3]byte{b[4], b[5], b[6]},
		DEOJ: [3]byte{b[7], b[8], b[9]},
		ESV:  b[10],
		OPC:  b[11],
	}

	pos := 12
	for i := 0; i < int(f.OPC); i++ {
		if pos+2 > len(b) {
			return nil, errors.New("frame too short in properties header")
		}

		pp := Property{
			EPC: b[pos],
			PDC: b[pos+1],
		}
		pos += 2

		if pos+int(pp.PDC) > len(b) {
			return nil, errors.New("frame too short in property data")
		}

		if pp.PDC > 0 {
			pp.EDT = append([]byte(nil), b[pos:pos+int(pp.PDC)]...)
		}

		pos += int(pp.PDC)

		f.Properties = append(f.Properties, pp)
	}

	return f, nil
}

func ParsePropertyMap(propertyMap Property) ([]byte, error) {
	if propertyMap.EPC != AnnouncePropertyMapEPC && propertyMap.EPC != SetPropertyMapEPC && propertyMap.EPC != GetPropertyMapEPC {
		return nil, errors.New("property map EPC is invalid")
	}

	if propertyMap.PDC == 0 {
		return nil, errors.New("property map PDC is empty")
	}

	var epcs []byte

	count := int(propertyMap.EDT[0])

	// プロパティマップ記述形式(1)
	if count < 16 && propertyMap.PDC == uint8(count+1) {
		for i := range count {
			epcs = append(epcs, propertyMap.EDT[1+i])
		}

		return epcs, nil
	}

	// プロパティマップ記述形式(2)
	if count >= 16 && propertyMap.PDC == uint8(17) {
		for i := range 16 {
			for bit := range 8 {
				if ((propertyMap.EDT[1+i] >> bit) & 0x01) == 1 {
					epcs = append(epcs, byte(0x8|bit)<<4|byte(i))
				}
			}
		}

		return epcs, nil
	}

	return nil, errors.New("invalid property map")
}
