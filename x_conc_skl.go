package main

import (
	"time"

	"github.com/benz9527/xboot/lib/list"
)

func xConcSklNew() list.SkipList[int, string] {
	skl, err := list.NewSkl[int, string](
		list.XConcSkl,
		func(i, j int) int64 {
			if i == j {
				return 0
			} else if i < j {
				return -1
			}
			return 1
		},
		list.WithXConcSklDataNodeUniqueMode[int, string](),
	)
	if err != nil {
		panic(err)
	}
	return skl
}

func xConcSklInserts(n int) {
	skl := xConcSklNew()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		skl.Insert(n-i, testString)
	}
}

func xConcSklWorstInserts(n int) {
	skl := xConcSklNew()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}
}

func xConcSklAvgSearch(n int) {
	skl := xConcSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.LoadFirst(i)
	}
}

func xConcSklSearchEnd(n int) {
	skl := xConcSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.LoadFirst(n)
	}
}

func xConcSklDelete(n int) {
	skl := xConcSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.RemoveFirst(i)
	}
}

func xConcSklWorstDelete(n int) {
	skl := xConcSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.RemoveFirst(n - i)
	}
}

var xConcSklFunctions = []func(int){
	xConcSklInserts,
	xConcSklWorstInserts,
	xConcSklAvgSearch,
	xConcSklSearchEnd,
	xConcSklDelete,
	xConcSklWorstDelete,
}
