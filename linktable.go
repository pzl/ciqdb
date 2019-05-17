package ciqdb

import (
	"encoding/binary"
	"strings"
)

type LinkTable struct {
	PRGSection
	links []Link
}

func (l *LinkTable) String() string {
	var buf strings.Builder
	buf.WriteString(l.PRGSection.String() + "\n")

	for _, link := range l.links {
		buf.WriteString("    " + link.String() + "\n")
	}

	return buf.String()
}

type Link struct {
	module int
	class  int
}

func (l Link) String() string {
	return "module: " + apidb(l.module) + ", class: " + apidb(l.class)
}

func parseDataTable(p *PRG, t SecType, length int, data []byte) *LinkTable {
	table := LinkTable{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
	}

	nLinks := int(binary.BigEndian.Uint16(data[0:2]))
	for i := 0; i < nLinks; i++ {
		table.links = append(table.links, Link{
			module: int(binary.BigEndian.Uint32(data[i*8+2 : i*8+6])),
			class:  int(binary.BigEndian.Uint32(data[i*8+6 : i*8+10])),
		})
	}

	return &table
}
