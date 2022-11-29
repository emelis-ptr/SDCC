package main

import (
	"SDCC-project/code/master-worker/master/master"
	"SDCC-project/code/util"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

/**
Benchmark
*/

var numWorker int
var numCentroid int
var numPoint int
var numMapper int
var numReducer int
var algo string

var regIP string
var regPort int
var masterPort int

func main() {
	var conf util.Conf
	conf.ReadConf(util.ConfPath)

	util.OpenEnv()
	numWorker, _ = strconv.Atoi(os.Getenv(util.NumWorker))
	numCentroid, _ = strconv.Atoi(os.Getenv(util.NumCluster))
	algo = os.Getenv(util.Algo)

	regIP = conf.RegIP
	regPort = conf.RegPort
	masterPort = conf.MasterPort

	res := testing.Benchmark(BenchmarkMaster20022)
	result(res)
	res = testing.Benchmark(BenchmarkMaster20052)
	result(res)
	res = testing.Benchmark(BenchmarkMaster50052)
	result(res)
	res = testing.Benchmark(BenchmarkMaster500105)
	result(res)
	res = testing.Benchmark(BenchmarkMaster8001010)
	result(res)
}

// result
func result(res testing.BenchmarkResult) {
	log.Println(" ************** ")
	log.Printf("Algorithm: %s", algo)
	log.Printf("Point: %d - Mapper: %d - Reducer: %d", numPoint, numMapper, numReducer)
	fmt.Printf("Memory allocations : %d \n", res.MemAllocs)
	fmt.Printf("Number of bytes processed in one iteration: %d \n", res.Bytes)
	fmt.Printf("Time taken (latency): %s \n", res.T)
	fmt.Printf("Number of bytes allocated: %d \n", res.MemBytes)
	log.Println(" ************** ")

	util.WriteBenchmark(res, numPoint, numMapper, numReducer, algo)
}

// BenchmarkMaster : benchmark
func benchmarkMaster(nPoint int, nMapper int, nReducer int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		master.Master(numWorker, nPoint, numCentroid, nMapper, nReducer, algo, regIP, regPort, masterPort, true)
	}
	numPoint = nPoint
	numMapper = nMapper
	numReducer = nReducer
}

func BenchmarkMaster20022(b *testing.B)   { benchmarkMaster(200, 2, 2, b) }
func BenchmarkMaster20052(b *testing.B)   { benchmarkMaster(200, 5, 2, b) }
func BenchmarkMaster50052(b *testing.B)   { benchmarkMaster(500, 5, 2, b) }
func BenchmarkMaster500105(b *testing.B)  { benchmarkMaster(500, 10, 5, b) }
func BenchmarkMaster8001010(b *testing.B) { benchmarkMaster(800, 10, 10, b) }
