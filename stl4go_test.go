package main

import (
	"testing"

	chen3feng "github.com/chen3feng/stl4go"
)

func BenchmarkChen3fengSkl(b *testing.B) {
	b.StopTimer()
	skl := chen3feng.NewSkipList[int, []byte]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		skl.Insert(i, testByteString)
	}
}
