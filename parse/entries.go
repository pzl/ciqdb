package main

import (
	"encoding/binary"
	"encoding/hex"
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
	symbol  int
	label   int
	icon    int
	apptype AppType
}

func (e *EntryPoint) String() string {
	return `
    UUID: ` + e.uuid + `
    Type: ` + e.apptype.String() + `
    ` + apidb(e.label) + `: ` + apidb(e.symbol) + `
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
		symbol:  int(binary.BigEndian.Uint32(data[20:24])),
		label:   int(binary.BigEndian.Uint32(data[24:28])),
		icon:    int(binary.BigEndian.Uint32(data[28:32])),
		apptype: AppType(binary.BigEndian.Uint32(data[32:36])),
	}
	// for symbol, label, also see .prg.debug : <symbolTable><entry id="X"/>

	return &entry, nil
}
