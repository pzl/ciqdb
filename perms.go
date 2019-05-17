package ciqdb

import (
	"encoding/binary"
	//"strconv"
	"strings"
)

type Permissions struct {
	PRGSection
	perms []Permission
}

func (p *Permissions) String() string {
	var buf strings.Builder
	buf.WriteString(p.PRGSection.String() + "\n")

	for _, perm := range p.perms {
		// so far, looks like you parse this as an int
		// and then match it with the entry in $SDK/bin/api.db
		// but there must be a SymbolTable way.. ?
		buf.WriteString("    " + perm.String() + "\n")
	}

	return buf.String()
}

type Permission struct {
	data []byte
}

func (p *Permission) String() string {
	return apidb(int(binary.BigEndian.Uint32(p.data)))
}

func parsePermissions(p *PRG, t SecType, length int, data []byte) *Permissions {
	perms := Permissions{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
	}

	numPerms := int(binary.BigEndian.Uint16(data[0:2]))
	for i := 0; i < numPerms; i++ {
		perms.perms = append(perms.perms, Permission{data[i*4+2 : i*4+6]})
	}

	return &perms
}
