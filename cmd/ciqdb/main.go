package main

import (
	"fmt"
	"os"

	"github.com/pzl/ciqdb"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("argument required: path to .prg file")
		os.Exit(1)
	}
	filename := os.Args[1]
	prg, err := ciqdb.NewPRG(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, s := range prg.Sections {
		fmt.Println(s)
	}
}
