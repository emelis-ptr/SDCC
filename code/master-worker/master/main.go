package main

import (
	"SDCC-project/code/master-worker/master/master"
	"SDCC-project/code/util"
	"fmt"
	"os"
	"strconv"
)

/*
*
Esecuzione del master
*/
func main() {
	fmt.Println("Master is up")

	var conf util.Conf
	conf.ReadConf(util.ConfPath)

	util.OpenEnv()
	numWorker, _ := strconv.Atoi(os.Getenv(util.NumWorker))
	numPoint, _ := strconv.Atoi(os.Getenv(util.NumPoint))
	numCentroid, _ := strconv.Atoi(os.Getenv(util.NumCluster))
	numMapper, _ := strconv.Atoi(os.Getenv(util.NumMapper))
	numReducer, _ := strconv.Atoi(os.Getenv(util.NumReducer))
	algo := os.Getenv(util.Algo)

	master.Master(numWorker, numPoint, numCentroid, numMapper, numReducer, algo, conf.RegIP, conf.RegPort, conf.MasterPort, false)
}
