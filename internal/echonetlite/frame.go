package echonetlite

import (
	"encoding/binary"
	"errors"
)

type Property struct {
	EPC byte
	PDC uint8
	EDT []byte
}

type Frame struct {
	EHD1       byte
	EHD2       byte
	TID        uint16
	SEOJ       [3]byte
	DEOJ       [3]byte
	ESV        byte
	OPC        uint8
	Properties []Property
}

func Parse(b []byte) (*Frame, error) {
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
		OPC:  uint8(b[11]),
	}

	pos := 12
	for i := 0; i < int(f.OPC); i++ {
		if pos+2 > len(b) {
			return nil, errors.New("frame too short in properties header")
		}

		pp := Property{
			EPC: b[pos],
			PDC: uint8(b[pos+1]),
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
