package ciqdb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

type DataSection struct {
	PRGSection
	data  []byte
	cdefs []ClassDef
}

func (d *DataSection) String() string {
	var buf strings.Builder

	buf.WriteString(d.PRGSection.String() + "\n")

	for _, c := range d.cdefs {
		buf.WriteString(c.String())
	}

	return buf.String()
}

type ClassDef struct {
	extendsOffset uint32
	staticEntry   uint32
	parentModule  uint32
	moduleID      uint32
	appTypes      uint16
	nFields       uint8 //so, a 255 limit
	fields        []Field
}

func (c ClassDef) String() string {
	var buf strings.Builder

	buf.WriteString("    extends Offset: " + strconv.FormatInt(int64(c.extendsOffset), 16) + "\n" +
		"    static Entry: " + strconv.FormatInt(int64(c.staticEntry), 16) + "\n" +
		"    parent module: " + apidb(int(c.parentModule)) + "\n" +
		"    module ID: " + apidb(int(c.moduleID)) + "\n" +
		"    app types: " + strconv.FormatInt(int64(c.appTypes), 16) + "\n" +
		"    fields:\n")

	for _, f := range c.fields {
		buf.WriteString("        " + f.String() + "\n")
	}
	return buf.String()
}

type Field struct {
	symbol    uint8
	flags     FieldFlag
	value     int
	valueType DataType
}

func (f Field) String() string {
	return apidb(int(f.symbol)) + ": " + strconv.Itoa(int(f.value)) + " " + f.flags.String() + " " + fmt.Sprint(f.valueType)
}

type FieldFlag uint8

const (
	FieldConst  FieldFlag = 1
	FieldHidden FieldFlag = 2
	FieldStatic FieldFlag = 4
)

func (f FieldFlag) String() string {
	var buf strings.Builder

	switch {
	case f&FieldConst != 0:
		buf.WriteString("const ")
	case f&FieldHidden != 0:
		buf.WriteString("hidden ")
	case f&FieldStatic != 0:
		buf.WriteString("static ")
	}
	return buf.String()
}

func (d *DataSection) getString(offset int) string {
	// I don't know what the first byte is for, but usually 0x01
	length := int(binary.BigEndian.Uint16(d.data[offset+1 : offset+3]))
	s := string(d.data[offset+3 : offset+3+length])
	return s
}

func parseData(p *PRG, t SecType, length int, data []byte) *DataSection {
	d := DataSection{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
		data: data,
	}
	classdef := []byte{0xC1, 0xA5, 0x5D, 0xEF}

	for i := 0; bytes.Equal(data[i:i+4], classdef); {
		cdef := ClassDef{
			extendsOffset: binary.BigEndian.Uint32(data[i+4 : i+8]),
			staticEntry:   binary.BigEndian.Uint32(data[i+8 : i+12]),
			parentModule:  binary.BigEndian.Uint32(data[i+12 : i+16]),
			moduleID:      binary.BigEndian.Uint32(data[i+16 : i+20]),
			appTypes:      binary.BigEndian.Uint16(data[i+20 : i+22]),
			nFields:       uint8(data[i+22]),
		}

		for j := 0; j < int(cdef.nFields); j++ {
			field := Field{
				symbol:    uint8(data[i+(j*8)+25]),
				flags:     FieldFlag(data[i+(j*8)+26] >> 4),
				value:     int(binary.BigEndian.Uint32(data[i+(j*8)+27 : i+(j*8)+31])),
				valueType: DataType(data[i+(j*8)+26] & 0x0F),
			}
			cdef.fields = append(cdef.fields, field)
		}
		i += 23 + 8*int(cdef.nFields)

		d.cdefs = append(d.cdefs, cdef)
	}

	return &d

}
