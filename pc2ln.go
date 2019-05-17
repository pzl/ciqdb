package ciqdb

import (
	"encoding/binary"
	"strconv"
	"strings"
)

type CodeTable struct {
	PRGSection
	Map []PCMap
}

func (c *CodeTable) String() string {
	var buf strings.Builder

	buf.WriteString(c.PRGSection.String() + "\n")

	for _, p := range c.Map {
		buf.WriteString("    " + p.String() + "\n")
	}

	return buf.String()
}

type PCMap struct {
	pc     int
	file   string
	symbol string
	line   int
}

func (p *PCMap) String() string {
	return p.file + ":" + strconv.Itoa(p.line) + " " + p.symbol + " (pc " + strconv.Itoa(p.pc) + ")"
}

func parsePCTable(p *PRG, t SecType, length int, data []byte) *CodeTable {
	table := CodeTable{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
	}

	var dataSec *DataSection
	for _, s := range p.Sections {
		if s.getType() == SectionData {
			dataSec = s.(*DataSection)
		}
	}

	if dataSec == nil {
		return nil
	}

	// skip first two data bytes. Count? Length?
	for i := 2; i < len(data); {
		thing := PCMap{
			pc:     int(binary.BigEndian.Uint32(data[i : i+4])),
			file:   dataSec.getString(int(binary.BigEndian.Uint32(data[i+4 : i+8]))),
			symbol: dataSec.getString(int(binary.BigEndian.Uint32(data[i+8 : i+12]))),
			line:   int(binary.BigEndian.Uint32(data[i+12 : i+16])),
		}
		i += 16
		table.Map = append(table.Map, thing)
	}

	return &table
}
