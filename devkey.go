package main

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
)

type DevKey struct {
	PRGSection
	signature []byte
	modulus   []byte
	exponent  int
}
func (d *DevKey) String() string {
	return d.PRGSection.String() + "\n" +
		"    signature: "+hex.EncodeToString(d.signature[0:10])+"...\n"+
		"    modulus: "+hex.EncodeToString(d.modulus[0:10])+"...\n"+
		"    exponent: "+strconv.Itoa(d.exponent)
}

func parseDevKey(p *PRG, t SecType, length int, data []byte) *DevKey {
	return &DevKey{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
		signature: data[0 : 512],
		modulus:   data[512 : 1024],
		exponent:  int(binary.BigEndian.Uint32(data[1024 : 1028])),
	}
}