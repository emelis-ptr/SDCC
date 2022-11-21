package main

import (
	"SDCC-project/code/master-worker/master/master"
	"SDCC-project/code/util"
	"fmt"
	"log"
	"testing"
)

/**
Benchmark
*/

var numPoint int
var numMapper int
var numReducer int
var algo string

func main() {
	for _, a := range algos {
		algo = a.input
		log.Printf("Algorithm: %s", a.input)
		for _, p := range points {
			numPoint = p.input
			for _, m := range mappers {
				numMapper = m.input
				if algo != util.KmeansPlusPlus {
					for _, r := range reducers {
						log.Printf("Point: %d - Mapper: %d - Reducer: %d", p.input, m.input, r.input)
						numReducer = r.input
						Benchmark()
					}
				} else {
					log.Printf("Point: %d - Mapper: %d - Reducer: %d", p.input, m.input, 1)
					Benchmark()
				}
			}
		}
	}
}

// Benchmark : esegue la funzione di benchmark
func Benchmark() {
	res := testing.Benchmark(BenchmarkMaster)
	fmt.Printf("Memory allocations : %d \n", res.MemAllocs)
	fmt.Printf("Number of bytes processed in one iteration: %d \n", res.Bytes)
	fmt.Printf("Number of run: %d \n", res.N)
	fmt.Printf("Time taken (latency): %s \n", res.T)
	fmt.Printf("Number of bytes allocated: %d \n", res.MemBytes)
	fmt.Printf("NsPerOp returns the ns/op metric: %d \n", res.NsPerOp())
	//AllocsPerOp returns the "allocs/op" metric, which is calculated as r.MemAllocs / r.N
	fmt.Printf("AllocsPerOp: %d \n", res.AllocsPerOp())
	//AllocedBytesPerOp returns the "B/op" metric, which is calculated as r.MemBytes / r.N
	fmt.Printf("AllocedBytesPerOp: %d \n", res.AllocedBytesPerOp())
}

// BenchmarkMaster : benchmark
func BenchmarkMaster(b *testing.B) {
	numCentroid := 5
	numWorker := 3
	regIP := "registry"
	regPort := 8000

	for n := 0; n < b.N; n++ {
		master.Master(numWorker, numPoint, numCentroid, numMapper, numReducer, algo, regIP, regPort, true)
	}
}

/** INPUT */

var points = []struct {
	input int
}{
	{input: 100},
	{input: 500},
	{input: 2000},
}

var mappers = []struct {
	input int
}{
	{input: 2},
	{input: 5},
	{input: 10},
}

var reducers = []struct {
	input int
}{
	{input: 2},
	{input: 5},
	{input: 10},
}

var algos = []struct {
	input string
}{
	{input: util.Llyod},
	{input: util.Standard},
	{input: util.KmeansPlusPlus},
}
