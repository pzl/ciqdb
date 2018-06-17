package main

import (
	"encoding/binary"
	"strconv"
	"strings"
)

type Exceptions struct {
	PRGSection
	exc []Exception
}
func (e *Exceptions) String() string {
	var buf strings.Builder

	buf.WriteString(e.PRGSection.String()+"\n")
	for _, ex := range e.exc {
		buf.WriteString("    Exception: try: " +
		                strconv.Itoa(ex.tryBegin) +
		                " - " + strconv.Itoa(ex.tryEnd) +
		                ". Handle: " + strconv.Itoa(ex.handleBegin))
	}
	return buf.String()
}

type Exception struct {
	tryBegin    int
	tryEnd      int
	handleBegin int
}

func i24ToInt(d []byte) int {
	return int(d[2]) | int(d[1]) << 8 | int(d[0]) << 16
}

func parseExceptions(p *PRG, t SecType, length int, data []byte) *Exceptions {
	e := Exceptions{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
	}

	count := int(binary.BigEndian.Uint16(data[0:2]))

	for i := 0; i < count; i++ {
		e.exc = append(e.exc, Exception{
			tryBegin: i24ToInt(data[i*9+2 : i*9+5]),
			tryEnd:   i24ToInt(data[i*9+5 : i*9+8]),
			handleBegin: i24ToInt(data[i*9+8 : i*9+11]),
		})
	}


	return &e
}
