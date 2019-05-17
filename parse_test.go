package ciqdb

import "testing"

func BenchmarkPRG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewPRG("samples/crystal-face.prg")
	}
}
