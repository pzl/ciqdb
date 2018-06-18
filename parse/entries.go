package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type EntrySection struct {
	PRGSection
	Entries []EntryPoint
}

func (e *EntrySection) String() string {
	var buf strings.Builder

	buf.WriteString(e.PRGSection.String())
	for i, ep := range e.Entries {
		buf.WriteString("\n  Entry Point " + strconv.Itoa(i))
		buf.WriteString(ep.String())
	}
	return buf.String()
}

type AppType uint8

const (
	WatchFace = AppType(iota)
	App
	DataField
	Widget
	Background
	Audio
)

func (a *AppType) String() string {
	name := []string{"Watch Face", "App", "Data Field", "Widget", "Background App", "Audio Provider"}
	switch {
	case *a <= Audio:
		return name[*a]
	default:
		return "Unknown (" + strconv.Itoa(int(*a)) + ")"
	}
}

type EntryPoint struct {
	uuid    string
	module  int
	symbol  string
	label   string
	icon    int
	apptype AppType
}

func (e *EntryPoint) String() string {
	return `
    UUID: ` + e.uuid + `
    Type: ` + e.apptype.String() + `
    ` + e.label + `: ` + e.symbol + `
    Module: ` + apidb(e.module) + `
    icon: ` + strconv.FormatInt(int64(e.icon), 16)
}

func parseEntries(p *PRG, t SecType, length int, data []byte) (*EntrySection, error) {
	e := EntrySection{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
	}

	n := int(binary.BigEndian.Uint16(data[:2]))
	for i := 0; i < n; i++ {
		entry, err := parseEntry(p.filename, data[i*36+2:(i+1)*36+2])
		if err != nil {
			return nil, err
		}
		e.Entries = append(e.Entries, *entry)
	}

	return &e, nil
}

func parseEntry(filename string, data []byte) (*EntryPoint, error) {
	entry := EntryPoint{
		uuid:    hex.EncodeToString(data[:16]),
		module:  int(binary.BigEndian.Uint32(data[16:20])),
		icon:    int(binary.BigEndian.Uint32(data[28:32])),
		apptype: AppType(binary.BigEndian.Uint32(data[32:36])),
	}

	symbol := int(binary.BigEndian.Uint32(data[20:24])) // see .prg.debug : <symbolTable><entry id="X"/>
	label := int(binary.BigEndian.Uint32(data[24:28]))

	var root struct {
		SymbolTable struct {
			Entry []struct {
				ID     int    `xml:"id,attr"`
				Module bool   `xml:"module,attr"`
				Field  bool   `xml:"field,attr"`
				Method bool   `xml:"method,attr"`
				Symbol string `xml:"symbol,attr"`
			} `xml:"entry"`
		} `xml:"symbolTable"`
	}

	file, err := os.Open(filename + ".debug.xml")
	if err == nil {
		defer file.Close()
		xmlData, e := ioutil.ReadAll(file)
		if e != nil {
			return &entry, e
		}
		xmlErr := xml.Unmarshal(xmlData, &root)
		if xmlErr != nil {
			return &entry, xmlErr
		}

		for _, s := range root.SymbolTable.Entry {
			if s.ID == symbol {
				entry.symbol = s.Symbol
			} else if s.ID == label {
				entry.label = s.Symbol
			}
		}
	}

	return &entry, nil
}
