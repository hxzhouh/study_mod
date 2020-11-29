package main

import "testing"

func BenchmarkInc(b *testing.B) {
	for i:=0;i<b.N;i++ {
		Inc()
	}
}
func BenchmarkIncRW(b *testing.B) {
	for i:=0;i<b.N;i++ {
		IncRW()
	}
}