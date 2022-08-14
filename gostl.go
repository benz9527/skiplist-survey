package main

import (
	"math"
	"time"

	liyue201 "github.com/liyue201/gostl/ds/skiplist"
)

func liyue201New(n int) *liyue201.Skiplist {
	// gostl's default level is only 10, so we must increase it to fit the test size
	return liyue201.New(liyue201.WithMaxLevel(int(math.Ceil(math.Log2(float64(n))))))
}

func liyue201Inserts(n int) {
	list := liyue201New(n)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}
}

func liyue201WorstInserts(n int) {
	list := liyue201New(n)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}
}

func liyue201AvgSearch(n int) {
	list := liyue201New(n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(i)
	}
}

func liyue201SearchEnd(n int) {
	list := liyue201New(n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(n)
	}
}

func liyue201Delete(n int) {
	list := liyue201New(n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(i)
	}
}

func liyue201WorstDelete(n int) {
	list := liyue201New(n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(n - i)
	}
}

var liyue201Functions = []func(int){liyue201Inserts, liyue201WorstInserts,
	liyue201AvgSearch, liyue201SearchEnd, liyue201Delete, liyue201WorstDelete}
