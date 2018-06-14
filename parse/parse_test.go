package main

import "testing"

func BenchmarkPRG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewPRG("/home/dan/dev/MonkeyC/samples/crystal-face/bin/crystal-face.prg")
	}
}
