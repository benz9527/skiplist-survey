package main

import (
	"time"

	chen3feng "github.com/chen3feng/stl4go"
)

func chen3fengInserts(n int) {
	list := chen3feng.NewSkipList[int, []byte]()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}
}

func chen3fengWorstInserts(n int) {
	list := chen3feng.NewSkipList[int, []byte]()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}
}

func chen3fengAvgSearch(n int) {
	list := chen3feng.NewSkipList[int, []byte]()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Find(i)
	}
}

func chen3fengSearchEnd(n int) {
	list := chen3feng.NewSkipList[int, []byte]()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Find(n)
	}
}

func chen3fengDelete(n int) {
	list := chen3feng.NewSkipList[int, []byte]()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(i)
	}
}

func chen3fengWorstDelete(n int) {
	list := chen3feng.NewSkipList[int, []byte]()

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(n - i)
	}
}

var chen3fengFunctions = []func(int){chen3fengInserts, chen3fengWorstInserts,
	chen3fengAvgSearch, chen3fengSearchEnd, chen3fengDelete, chen3fengWorstDelete}
