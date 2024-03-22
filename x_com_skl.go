package main

import (
	"time"

	"github.com/benz9527/xboot/lib/list"
)

var (
	testString = string(testByteString)
)

func xComSklNew() list.SkipList[int, string] {
	skl, err := list.NewSkl[int, string](
		list.XComSkl,
		func(i, j int) int64 {
			if i == j {
				return 0
			} else if i < j {
				return -1
			}
			return 1
		},
	)
	if err != nil {
		panic(err)
	}
	return skl
}

func xComSklInserts(n int) {
	skl := xComSklNew()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		skl.Insert(n-i, testString)
	}
}

func xComSklWorstInserts(n int) {
	skl := xComSklNew()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}
}

func xComSklAvgSearch(n int) {
	skl := xComSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.LoadFirst(i)
	}
}

func xComSklSearchEnd(n int) {
	skl := xComSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.LoadFirst(n)
	}
}

func xComSklDelete(n int) {
	skl := xComSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.RemoveFirst(i)
	}
}

func xComSklWorstDelete(n int) {
	skl := xComSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.RemoveFirst(n - i)
	}
}

var xComSklFunctions = []func(int){
	xComSklInserts,
	xComSklWorstInserts,
	xComSklAvgSearch,
	xComSklSearchEnd,
	xComSklDelete,
	xComSklWorstDelete,
}
