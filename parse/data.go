package main

import (
	"encoding/binary"
)

type DataSection struct {
	PRGSection
	data []byte
}


func (d *DataSection) getString(offset int) string {
	// I don't know what the first byte is for, but usually 0x01
	length := int(binary.BigEndian.Uint16(d.data[offset+1:offset+3]))
	s := string(d.data[offset+3 : offset+3+length])
	return s
}

func parseData(p *PRG, t SecType, length int, data []byte) *DataSection {
	return &DataSection {
		PRGSection: PRGSection {
			Type:   t,
			length: length,
		},
		data: data,
	}
}

/*



*/