package kmeansAlgo

import (
	"main/code/mapreduce"
	"math/rand"
	"net/rpc"
)

/**
1. Choose k initial means µ1, . . . , µk uniformly at random from the set X.
2. Apply the MapReduce given by k-meansMap and k-meansReduce to X.
3. Compute the new means µ1, . . . , µk from the results of the MapReduce.
4. Broadcast the new means to each machine on the cluster.
5. Repeat steps 2 through 4 until the means have converged.
*/

func InitCentroidStandard(points []mapreduce.Points, numCentroid int) []mapreduce.Centroids {
	centroids := make([]mapreduce.Centroids, numCentroid)

	for i := 0; i < numCentroid; i++ {
		min := i * len(points) / numCentroid
		max := ((i + 1) * len(points)) / numCentroid
		centroids[i].Index = i
		centroids[i].Centroid = points[min:max][rand.Intn(len(points[min:max]))].Point
	}

	return centroids
}

func StandardKMeans(numWorker int, numCentroid int, points []mapreduce.Points, algo string, clients []*rpc.Client, calls []*rpc.Call) {
	//Iterazione uguale all'algoritmo di Llyod
	Llyod(numWorker, numCentroid, points, algo, clients, calls)
}
