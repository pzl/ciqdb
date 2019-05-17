package ciqdb

import (
	"fmt"
	"io"
	"os"
)

//go:generate stringer -type=DataType
type DataType uint8

const (
	Null DataType = iota
	Int
	Float
	String
	Object
	Array
	Method
	ClassDefinition
	Symbol
	Bool
	ModuleDef
	Hash
	Resource // then resourceBITMAP=0, FONT=1
	PrimitiveObj
	Long
	Double
	WeakPointer
	PrimitiveMod
	SysPointer
	Char
)

type Section interface {
	getLength() int
	getName() string
	getLabel() string
	getType() SecType
	fmt.Stringer // String() string
}

type PRG struct {
	Filename string
	Sections []Section
}

func (p *PRG) Parse(r io.Reader) error {
	for {
		s, err := readSection(p, r)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		p.Sections = append(p.Sections, s)
	}
	return nil
}

func NewPRG(filename string) (*PRG, error) {
	prg := PRG{}
	prg.Filename = filename

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := prg.Parse(file); err != nil {
		return nil, err
	}

	return &prg, nil
}
