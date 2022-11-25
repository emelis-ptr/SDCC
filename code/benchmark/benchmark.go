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
				for _, r := range util.Reducers {
					numReducer = r.Input
					Benchmark()
				}
			}
		}
	}
}

// Benchmark : esegue la funzione di benchmark
func Benchmark() {
	res := testing.Benchmark(BenchmarkMaster)
	log.Println(" ************** ")
	log.Printf("Point: %d - Mapper: %d - Reducer: %d", numPoint, numMapper, numReducer)
	fmt.Printf("Memory allocations : %d \n", res.MemAllocs)
	fmt.Printf("Number of bytes processed in one iteration: %d \n", res.Bytes)
	fmt.Printf("Time taken (latency): %s \n", res.T)
	fmt.Printf("Number of bytes allocated: %d \n", res.MemBytes)
	log.Println(" ************** ")

	util.WriteBenchmark(res, numPoint, numMapper, numReducer, algo)
}

// BenchmarkMaster : benchmark
func BenchmarkMaster(b *testing.B) {
	numCentroid := 5
	numWorker := 3
	regIP := "registry"
	regPort := 8000
	masterPort := 9000

	for n := 0; n < b.N; n++ {
		master.Master(numWorker, numPoint, numCentroid, numMapper, numReducer, algo, regIP, regPort, masterPort, true)
	}
}
