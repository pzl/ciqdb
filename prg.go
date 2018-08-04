package main

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
	filename string
	sections []Section
}

func (p *PRG) Parse(r io.Reader) error {
	for {
		s, err := readSection(p, r)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		p.sections = append(p.sections, s)
	}
	return nil
}

func NewPRG(filename string) (*PRG, error) {
	prg := PRG{}
	prg.filename = filename

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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("argument required: path to .prg file")
		os.Exit(1)
	}
	filename := os.Args[1]
	prg, err := NewPRG(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, s := range prg.sections {
		fmt.Println(s)
	}
}
