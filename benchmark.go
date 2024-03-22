package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	// The range of values for n passed to the individual benchmarks
	start, end, step int
)

var (
	wg             sync.WaitGroup
	testByteString = []byte(fmt.Sprint("test value"))
	test10Kilobyte = make([]byte, 10240)
)

func init() {
	flag.IntVar(&start, "flagname", 100000, "help message for flagname")
	flag.IntVar(&end, "end", 3000000, "help message for flagname")
	flag.IntVar(&step, "step", 100000, "help message for flagname")
}

var filter = flag.String("filter", ".*", "comma separated regexps to filter run test cases")

// timeTrack will print out the number of nanoseconds since the start time divided by n
// Useful for printing out how long each iteration took in a benchmark
func timeTrack(start time.Time, n int) {
	loopNS := time.Since(start).Nanoseconds() / int64(n)
	fmt.Print(loopNS)
}

// iterations is used to print out the CSV header with iteration counts
func iterations(n int) {
	fmt.Print(n)
}

// funcName returns just the function name of a string, given any function at all
func funcName(f func(int)) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()[5:]
}

// runIterations executes the tests in a loop with the given parameters
func runIterations(name string, start, end, step int, f func(int)) {
	fmt.Print(name, ",")
	for i := start; i <= end; i += step {
		f(i)
		fmt.Print(",")
	}
	fmt.Println()
}

func matchFuncName(filters []string, name string) bool {
	for _, f := range filters {
		m, err := regexp.MatchString(f, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid filter '%v'", f)
			os.Exit(1)
		}
		if m {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()
	// first, print the CSV header with iteration counts
	runIterations("iterations", start, end, step, iterations)
	var allFunctions []func(int)
	allFunctions = append(allFunctions, xComSklFunctions...)
	allFunctions = append(allFunctions, xConcSklFunctions...)

	allFunctions = append(allFunctions, chen3fengFunctions...)
	allFunctions = append(allFunctions, colFunctions...)
	allFunctions = append(allFunctions, huanduFunctions...)
	allFunctions = append(allFunctions, liyue201Functions...)
	allFunctions = append(allFunctions, mtchavezFunctions...)
	allFunctions = append(allFunctions, mtFunctions...)
	allFunctions = append(allFunctions, seanFunctions...)
	allFunctions = append(allFunctions, ryszardFunctions...)

	filters := strings.Split(*filter, ",")
	for _, f := range allFunctions {
		fname := funcName(f)
		if matchFuncName(filters, fname) {
			runIterations(fname, start, end, step, f)
		}
	}
}
