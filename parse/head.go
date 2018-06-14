package main

import (
	"encoding/binary"
	"strconv"
)

/*
	D000D000 label
	0000000D length (13 - possibly static)
	00 - constant
	02 04 05 - API version
	00 00 00 00 mBackgroundDataOffset
	00 00 00 00 mBackgroundCodeOffset
	00 mEnableAppTrial

*/

type Version struct {
	major uint8
	minor uint8
	patch uint8
}

func (v *Version) String() string {
	return strconv.Itoa(int(v.major)) + "." + strconv.Itoa(int(v.minor)) + "." + strconv.Itoa(int(v.patch))
}

type Head struct {
	PRGSection
	CIQVersion   Version
	BGDataOffset int
	BGCodeOffset int
	AppTrial     bool
}

func (h *Head) String() string {
	return h.PRGSection.String() + `
    CIQ version: ` + h.CIQVersion.String() + `
    App Trial Enabled: ` + strconv.FormatBool(h.AppTrial)
}

func parseHead(t SecType, length int, data []byte) *Head {
	return &Head{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
		CIQVersion:   Version{data[1], data[2], data[3]},
		BGDataOffset: int(binary.BigEndian.Uint32(data[4:8])),
		BGCodeOffset: int(binary.BigEndian.Uint32(data[8:12])),
		AppTrial:     data[12] == 1,
	}
}
