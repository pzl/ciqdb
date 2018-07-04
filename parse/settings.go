package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type SettingVal struct {
	datatype DataType
	offset   int
	value    []byte
}

func (s *SettingVal) String() string {
	switch s.datatype {
	case Bool:
		return strconv.FormatBool(int(s.value[0]) == 0)
	case Int:
		return strconv.Itoa(int(binary.BigEndian.Uint32(s.value)))
	case String:
		return string(s.value)
	// todo: float
	default:
		return hex.EncodeToString(s.value)
	}
}

type Settings map[string]SettingVal

type SettingsSection struct {
	PRGSection
	S Settings
}

func (s *SettingsSection) String() string {
	var str strings.Builder
	str.WriteString(s.PRGSection.String() + "\n")
	for k, v := range s.S {
		str.WriteString("    " + k + ": " + v.String() + "\n")
	}
	return str.String()
}

func parseSettings(p *PRG, t SecType, length int, data []byte) *SettingsSection {
	s := SettingsSection{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
		S: Settings{},
	}
	var strs map[int]string
	vals := []SettingVal{}
	for i := 0; i < len(data); {
		subsec := hex.EncodeToString(data[i : i+4])
		i += 4
		sublen := int(binary.BigEndian.Uint32(data[i : i+4]))
		i += 4
		if subsec == "abcdabcd" {
			strs = parseSettingsStrings(data[i : i+sublen])
		} else if subsec == "da7ada7a" {
			vals = parseSettingsValues(data[i : i+sublen])
		} else {
			fmt.Println("unknown settings section found: " + subsec + " (length " + strconv.Itoa(sublen) + ")")
		}
		i += sublen
	}

	for _, v := range vals {
		if v.datatype == String {
			v.value = []byte(strs[int(binary.BigEndian.Uint32(v.value))])
		}
		s.S[strs[v.offset]] = v
	}
	return &s
}

func parseSettingsStrings(data []byte) map[int]string {
	names := map[int]string{}
	for i := 0; i < len(data); {
		length := int(binary.BigEndian.Uint16(data[i : i+2]))
		name := string(data[i+2 : i+2+length-1]) //-1 to not include NUL byte
		names[i] = name
		i += 2 + length
	}
	return names
}

func parseSettingsValues(data []byte) []SettingVal {
	vals := []SettingVal{}

	// first byte is 0x0B? significance unknown
	// next 4-byte is the number of entries
	// skip these first 5 bytes

	for i := 5; i < len(data); {
		i += 1 //skip start byte, should be constant 0x03
		offset := int(binary.BigEndian.Uint32(data[i : i+4]))
		dt := DataType(data[i+4])
		var n int
		switch dt {
		case Bool:
			n = 1
		default:
			n = 4
		}
		value := data[i+5 : i+5+n]
		vals = append(vals, SettingVal{dt, offset, value})
		i += 5 + n
	}
	return vals
}
