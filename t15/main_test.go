package main

import "testing"

func BenchmarkHugeStrCut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hugeStrCut()
	}
}

func BenchmarkHugeSliceCut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hugeSliceCut()
	}
}

func BenchmarkHugeStrRet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hugeStrRet()
	}
}

func BenchmarkHugeRuneRet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hugeRuneRet()
	}
}

func BenchmarkSmallStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallStr()
	}
}

func BenchmarkSmallRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		smallRune()
	}
}

// go test -bench . ./... -benchmem
