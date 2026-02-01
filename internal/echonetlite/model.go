package echonetlite

const (
	VersionEPC             = 0x82
	AnnouncePropertyMapEPC = 0x9d
	SetPropertyMapEPC      = 0x9e
	GetPropertyMapEPC      = 0x9f
)

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

type Property struct {
	EPC byte
	PDC uint8
	EDT []byte
}

type ClassDefinition struct {
	IsNodeProfile bool
	Name          string
	NameEN        string
	ShortName     string
	Properties    []PropertyDefinition
}

type PropertyDefinition struct {
	EPC           byte
	Name          string
	NameEN        string
	ShortName     string
	Description   string
	DescriptionEN string
	ValidRelease  struct {
		FROM byte
		TO   byte
	}
}

var ClassDefinitions map[[2]byte]ClassDefinition
