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
	for _, a := range util.Algos {
		algo = a.Input
		log.Printf("Algorithm: %s", a.Input)
		for _, p := range util.Points {
			numPoint = p.Input
			for _, m := range util.Mappers {
				numMapper = m.Input
				if algo != util.KmeansPlusPlus {
					for _, r := range util.Reducers {
						log.Printf("Point: %d - Mapper: %d - Reducer: %d", p.Input, m.Input, r.Input)
						numReducer = r.Input
						Benchmark()
					}
				} else {
					log.Printf("Point: %d - Mapper: %d - Reducer: %d", p.Input, m.Input, 1)
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
