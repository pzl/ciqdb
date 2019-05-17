package ciqdb

var SymTable *SymbolTable

func apidb(x int) string {
	if SymTable == nil {
		return "[no symbol table]"
	}
	return SymTable.Lookup(x)
}
