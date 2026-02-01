package echonetlite_test

import (
	"testing"

	"github.com/lrks/kodama-net/internal/echonetlite"
)

func TestParseFrame_ValidFrame(t *testing.T) {
	b := []byte{
		0x10, 0x81, // EHD1, EHD2
		0x12, 0x34, // TID
		0x05, 0xff, 0x01, // SEOJ
		0x0e, 0xf0, 0x01, // DEOJ
		0x62,       // ESV
		0x01,       // OPC
		0xd6,       // EPC
		0x02,       // PDC
		0x01, 0x02, // EDT
	}

	f, err := echonetlite.ParseFrame(b)
	if err != nil {
		t.Fatalf("ParseFrame returned error: %v", err)
	}

	if f.EHD1 != 0x10 || f.EHD2 != 0x81 {
		t.Fatalf("unexpected header: %x %x", f.EHD1, f.EHD2)
	}

	if f.TID != 0x1234 {
		t.Fatalf("unexpected TID: %04x", f.TID)
	}

	if len(f.Properties) != 1 {
		t.Fatalf("expected 1 property, got %d", len(f.Properties))
	}

	p := f.Properties[0]
	if p.EPC != 0xd6 || p.PDC != 2 {
		t.Fatalf("unexpected property header: EPC=%x PDC=%d", p.EPC, p.PDC)
	}

	if len(p.EDT) != 2 || p.EDT[0] != 0x01 || p.EDT[1] != 0x02 {
		t.Fatalf("unexpected EDT: %v", p.EDT)
	}
}

func TestParseFrame_ErrTooShort(t *testing.T) {
	_, err := echonetlite.ParseFrame([]byte{0x10, 0x81})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestParseFrame_UnsupportedHeader(t *testing.T) {
	b := make([]byte, 12)
	b[0] = 0x11 // wrong EHD1
	b[1] = 0x81

	_, err := echonetlite.ParseFrame(b)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
