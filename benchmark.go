package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	// The range of values for n passed to the individual benchmarks
	start, end, step int
)

var (
	wg             sync.WaitGroup
	testByteString = []byte(fmt.Sprint("test value"))
	test10Kilobyte = make([]byte, 10240)
	dataC          = make(chan int64, 1)
	chartCols      = make([]string, 0, 16)
	insertMap      = make(map[string][]int64)
	worstInsertMap = make(map[string][]int64)
	avgSearchMap   = make(map[string][]int64)
	searchEndMap   = make(map[string][]int64)
	deleteMap      = make(map[string][]int64)
	worstDeleteMap = make(map[string][]int64)
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
	dataC <- loopNS
}

// funcName returns just the function name of a string, given any function at all
func funcName(f func(int)) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()[5:]
}

// runIterations executes the tests in a loop with the given parameters
func runIterations(fname string, start, end, step int, f func(int)) {
	var arr = make([]int64, 0, 16)
	for i := start; i <= end; i += step {
		f(i)
		data := <-dataC
		arr = append(arr, data)
	}
	if !strings.HasSuffix(fname, "WorstInserts") && strings.HasSuffix(fname, "Inserts") {
		insertMap[fname] = arr
	} else if strings.HasSuffix(fname, "WorstInserts") {
		worstInsertMap[fname] = arr
	} else if strings.HasSuffix(fname, "AvgSearch") {
		avgSearchMap[fname] = arr
	} else if strings.HasSuffix(fname, "SearchEnd") {
		searchEndMap[fname] = arr
	} else if !strings.HasSuffix(fname, "WorstDelete") && strings.HasSuffix(fname, "Delete") {
		deleteMap[fname] = arr
	} else if strings.HasSuffix(fname, "WorstDelete") {
		worstDeleteMap[fname] = arr
	}
}

func genIterationsCols(start, end, step int) {
	for i := start; i <= end; i += step {
		chartCols = append(chartCols, strconv.Itoa(i))
	}
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
	genIterationsCols(start, end, step)
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

	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Skl bench",
		Subtitle: "Inserts",
	}))
	// Put data into instance
	line.SetXAxis(chartCols[:])
	for k, v := range insertMap {
		items := make([]opts.LineData, 0, len(v))
		for i := 0; i < len(v); i++ {
			items = append(items, opts.LineData{Value: v[i]})
		}
		line.AddSeries(k, items)
	}
	path, err := filepath.Abs("./bench/inserts.html")
	if err != nil {
		panic(err)
	}
	f, _ := os.Create(path)
	line.Render(f)
	f.Close()

	line = charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Skl bench",
		Subtitle: "WorstInserts",
	}))
	// Put data into instance
	line.SetXAxis(chartCols[:])
	for k, v := range worstInsertMap {
		items := make([]opts.LineData, 0, len(v))
		for i := 0; i < len(v); i++ {
			items = append(items, opts.LineData{Value: v[i]})
		}
		line.AddSeries(k, items)
	}
	path, err = filepath.Abs("./bench/worst-inserts.html")
	if err != nil {
		panic(err)
	}
	f, _ = os.Create(path)
	line.Render(f)
	f.Close()

	line = charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Skl bench",
		Subtitle: "AvgSearch",
	}))
	line.SetXAxis(chartCols[:])
	// Put data into instance
	line.SetXAxis(chartCols[:])
	for k, v := range avgSearchMap {
		items := make([]opts.LineData, 0, len(v))
		for i := 0; i < len(v); i++ {
			items = append(items, opts.LineData{Value: v[i]})
		}
		line.AddSeries(k, items)
	}
	path, err = filepath.Abs("./bench/avg-search.html")
	if err != nil {
		panic(err)
	}
	f, _ = os.Create(path)
	line.Render(f)
	f.Close()

	line = charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Skl bench",
		Subtitle: "SearchEnd",
	}))
	line.SetXAxis(chartCols[:])
	for k, v := range searchEndMap {
		items := make([]opts.LineData, 0, len(v))
		for i := 0; i < len(v); i++ {
			items = append(items, opts.LineData{Value: v[i]})
		}
		line.AddSeries(k, items)
	}
	path, err = filepath.Abs("./bench/search-end.html")
	if err != nil {
		panic(err)
	}
	f, _ = os.Create(path)
	line.Render(f)
	f.Close()

	line = charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Skl bench",
		Subtitle: "Deletes",
	}))
	// Put data into instance
	line.SetXAxis(chartCols[:])
	for k, v := range deleteMap {
		items := make([]opts.LineData, 0, len(v))
		for i := 0; i < len(v); i++ {
			items = append(items, opts.LineData{Value: v[i]})
		}
		line.AddSeries(k, items)
	}
	path, err = filepath.Abs("./bench/deletes.html")
	if err != nil {
		panic(err)
	}
	f, _ = os.Create(path)
	line.Render(f)
	f.Close()

	line = charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Skl bench",
		Subtitle: "WorstDeletes",
	}))
	// Put data into instance
	line.SetXAxis(chartCols[:])
	for k, v := range worstDeleteMap {
		items := make([]opts.LineData, 0, len(v))
		for i := 0; i < len(v); i++ {
			items = append(items, opts.LineData{Value: v[i]})
		}
		line.AddSeries(k, items)
	}
	path, err = filepath.Abs("./bench/worst-deletes.html")
	if err != nil {
		panic(err)
	}
	f, _ = os.Create(path)
	line.Render(f)
	f.Close()

}
