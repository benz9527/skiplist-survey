package main

import (
	"time"

	"github.com/benz9527/xboot/lib/list"
)

func xConcSklNew() list.SkipList[int, []byte] {
	skl, err := list.NewSkl[int, []byte](
		list.XConcSkl,
		func(i, j int) int64 {
			if i == j {
				return 0
			} else if i < j {
				return -1
			}
			return 1
		},
		list.WithXConcSklDataNodeUniqueMode[int, []byte](true),
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
		skl.Insert(n-i, testByteString)
	}
}

func xConcSklWorstInserts(n int) {
	skl := xConcSklNew()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
	}
}

func xConcSklAvgSearch(n int) {
	skl := xConcSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.LoadFirst(i)
	}
}

func xConcSklSearchEnd(n int) {
	skl := xConcSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.LoadFirst(n)
	}
}

func xConcSklDelete(n int) {
	skl := xConcSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.RemoveFirst(i)
	}
}

func xConcSklWorstDelete(n int) {
	skl := xConcSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
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
