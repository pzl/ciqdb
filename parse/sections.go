package main

import (
	"encoding/binary"
	"io"
	"strconv"
	"strings"
)

type PRGSection struct {
	Type   SecType
	length int
}

func (s *PRGSection) String() string {
	return s.getLabel() + " - " + s.getName() + " (" + strconv.Itoa(s.getLength()) + " bytes)"
}
func (s *PRGSection) getName() string {
	return s.Type.String()
}
func (s *PRGSection) getLabel() string {
	label := strconv.FormatInt(int64(s.Type), 16)
	if len(label) < 8 {
		return strings.Repeat("0", 8-len(label)) + label
	}
	return label
}
func (s *PRGSection) getLength() int {
	return s.length
}
func (s *PRGSection) getType() SecType {
	return s.Type
}

type GenericDataSection struct {
	PRGSection
	data []byte
}

type SecType int

const (
	SectionHead       = SecType(0xd000d000)
	SectionEntry      = SecType(0x6060c0de) // "gogo code"
	SectionData       = SecType(0xda7ababe)
	SectionCode       = SecType(0xc0debabe)
	SectionCodeTable  = SecType(0xc0de7ab1)
	SectionClassTable = SecType(0xc1a557b1)
	SectionFoodGood   = SecType(0xf00d600d)
	SectionGoodBoy    = SecType(0x6000db01)
	SectionExceptions = SecType(0x0ece7105) // ecetios
	SectionSymbols    = SecType(0x5717b015) // stltbols?
	SectionSettings   = SecType(0x5e771465) // settings
	SectionEncoded    = SecType(0xe1c0de12) // elcodel
	SectionUnlock     = SecType(0xd011aaa5) // unlock for app trials?
	SectionIQSig      = SecType(0x00020833)
	SectionEnd        = SecType(0x00000000)
)

var secNames = map[SecType]string{
	SectionHead:       "Head",
	SectionEntry:      "Entry Points",
	SectionData:       "Data",
	SectionCode:       "Code",
	SectionCodeTable:  "Code Table (aka PCtoLineNum)",
	SectionClassTable: "Class Table (class imports)",
	SectionFoodGood:   "Resources",
	SectionGoodBoy:    "Permissions",
	SectionExceptions: "Exceptions",
	SectionSymbols:    "Symbols",
	SectionSettings:   "Settings",
	SectionEncoded:    "Developer Signature",
	SectionUnlock:     "App Unlock",
	SectionIQSig:      "App Store Signature",
	SectionEnd:        "End",
}

func (t SecType) String() string {
	if name, exists := secNames[t]; exists {
		return name
	} else {
		return "Unknown"
	}
}

func readSection(p *PRG, f io.Reader) (Section, error) {
	head := make([]byte, 8)
	if _, err := io.ReadFull(f, head); err != nil {
		return nil, err
	}

	secType := SecType(binary.BigEndian.Uint32(head[:4]))
	secLength := int(binary.BigEndian.Uint32(head[4:]))

	//to skip reading it in
	// f.Seek(int64(sec.length), 1)
	data := make([]byte, secLength)
	if _, err := io.ReadFull(f, data); err != nil {
		return nil, err
	}

	var sec Section
	var err error
	switch secType {
	case SectionHead:
		sec = parseHead(secType, secLength, data)
	case SectionEntry:
		sec, err = parseEntries(p, secType, secLength, data)
	case SectionSettings:
		sec = parseSettings(p, secType, secLength, data)
	case SectionData:
		sec = parseData(p, secType, secLength, data)
	case SectionCodeTable:
		sec = parsePCTable(p, secType, secLength, data)
	case SectionGoodBoy:
		sec = parsePermissions(p, secType, secLength, data)
	default:
		sec = defaultSection(secType, secLength, data)
	}
	if err != nil {
		return nil, err
	}
	return sec, nil
}

func defaultSection(t SecType, length int, data []byte) *GenericDataSection {
	return &GenericDataSection{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
		data: data,
	}
}
