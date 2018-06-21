package main

import (
	"encoding/binary"
	"strconv"
	"strings"
)

type SymbolTable struct {
	PRGSection
	table map[int]string
}

func (st *SymbolTable) String() string {
	var buf strings.Builder

	buf.WriteString(st.PRGSection.String() + "\n")

	for k, v := range st.table {
		buf.WriteString("    " + strconv.Itoa(k) + ": " + v + "\n")
	}

	return buf.String()
}

func (st *SymbolTable) Lookup(id int) string {
	if val, exists := st.table[id]; exists {
		return val
	}
	return "Unknown Symbol ID: "+strconv.Itoa(id)
}

func parseSymbols(p *PRG, t SecType, length int, data []byte) *SymbolTable {
	st := SymbolTable{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
		table: make(map[int]string),
	}

	n := int(binary.BigEndian.Uint16(data[0:2]))

	for i := 0; i < n; i++ {
		id := int(binary.BigEndian.Uint32(data[2+i*8 : 6+i*8]))
		offset := int(binary.BigEndian.Uint32(data[6+i*8 : 10+i*8]))

		// first byte is 0x01, String type, I think
		slen := int(binary.BigEndian.Uint16(data[offset+1 : offset+3]))
		s := string(data[offset+3 : offset+3+slen])
		st.table[id] = s
	}

	SymTable = &st

	return &st
}
