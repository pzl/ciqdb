package main

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

	for _,p := range c.Map {
		buf.WriteString("    " + p.String()+"\n")
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
	for _, s := range p.sections {
		if s.getType() == SectionData {
			dataSec = s.(*DataSection)
		}
	}

	if dataSec == nil {
		return nil
	}

	// skip first two data bytes. Count? Length?
	for i:=2; i < len(data); {
		thing := PCMap{
			pc: int(binary.BigEndian.Uint32(data[i:i+4])),
			file: dataSec.getString(int(binary.BigEndian.Uint32(data[i+4:i+8]))),
			symbol: dataSec.getString(int(binary.BigEndian.Uint32(data[i+8 : i+12]))),
			line: int(binary.BigEndian.Uint32(data[i+12:i+16])),			
		}
		i+=16
		table.Map = append(table.Map, thing)
	}



	/*
	@TODO:

	data looks like this:
03 16 (possible count? length?)

10 00 00 00 = pc = 268435456
00 00 0D D8 = FILE = 3544
00 00 0E 70 = SYMBOL = 3696
00 00 00 23 = lineNum = 35


10 00 00 04 = pc = 268435456
00 00 0D D8
00 00 0E 70 = 3696
00 00 00 24 = lineNum = 36


10 00 00 18 = pc
00 00 0D D8
00 00 0E 70 = 3696
00 00 00 26 = lineNum = 38

10 00 00 28 = pc
00 00 0D D8
00 00 0E 70 = 3696
00 00 00 27 = lineNum = 39

...

10 00 00 68 = pc = 268435560
00 00 0D D8
00 00 0E 70 = 3696
00 00 00 2C = 44


10 00 00 7C = pc = 268435580
00 00 0D D8
00 00 0C 80 = 3200
00 00 00 2F = 47

	pc and lineNum are just easy int32's
	FILE and SYMBOL are strings looked up via PRG's Data section, as an offset.
	
	the strings have 2-byte length, then the string starts

	*/
	return &table
}
