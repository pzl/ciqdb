package main

var SymTable *SymbolTable

func apidb(x int) string {
	return SymTable.Lookup(x)
}
