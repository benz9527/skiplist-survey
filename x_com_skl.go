package main

import (
	"time"

	"github.com/benz9527/xboot/lib/list"
)

func xComSklNew() list.SkipList[int, []byte] {
	skl, err := list.NewSkl[int, []byte](list.XComSkl)
	if err != nil {
		panic(err)
	}
	return skl
}

func xComSklInserts(n int) {
	skl := xComSklNew()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		skl.Insert(n-i, testByteString)
	}
}

func xComSklWorstInserts(n int) {
	skl := xComSklNew()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
	}
}

func xComSklAvgSearch(n int) {
	skl := xComSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.LoadFirst(i)
	}
}

func xComSklSearchEnd(n int) {
	skl := xComSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.LoadFirst(n)
	}
}

func xComSklDelete(n int) {
	skl := xComSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = skl.RemoveFirst(i)
	}
}

func xComSklWorstDelete(n int) {
	skl := xComSklNew()

	for i := 0; i < n; i++ {
		skl.Insert(i, testByteString)
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
