package util

const (
	NumWorker  = "NUMWORKER"
	NumPoint   = "NUMPOINT"
	NumCluster = "NUMCLUSTER"
	Algo       = "ALGO"
	NumVector  = 2

	NumMapper  = "NUMMAPPER"
	NumReducer = "NUMREDUCER"

	Llyod          = "llyod"
	Standard       = "standardKMeans"
	KmeansPlusPlus = "algorithm"

	NameFile = "algo"

	ConfPath = "./config.json"
)

/** INPUT */

var Points = []struct {
	Input int
}{
	{Input: 100},
	{Input: 500},
	{Input: 10000},
}

var Mappers = []struct {
	Input int
}{
	{Input: 2},
	{Input: 5},
	{Input: 10},
}

var Reducers = []struct {
	Input int
}{
	{Input: 2},
	{Input: 5},
	{Input: 10},
}

var Algos = []struct {
	Input string
}{
	{Input: Llyod},
	{Input: Standard},
	{Input: KmeansPlusPlus},
}
