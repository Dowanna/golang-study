package main

import (
	"testing"
)

var testArray = []string{"hogehoge", "fugafuga", "piyopiyo"}

func BenchmarkForSlow(b *testing.B) {
	b.ResetTimer()
	slow(testArray)
}

func BenchmarkForFast(b *testing.B) {
	b.ResetTimer()
	fast(testArray)
}
