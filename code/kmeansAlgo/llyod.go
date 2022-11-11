package kmeansAlgo

import (
	"log"
	"main/code/mapreduce"
	"math/rand"
	"net/rpc"
)

/**
Specify desired number of clusters k;
Initially choose k data points that are likely to be in different clusters;
Make these data points the centroids of their clusters;
Repeat
For each remaining data point p do
Find the centroid to which p is nearest;
Add p to the cluster of that centroid;
Re-compute cluster centroids;
Until no improvement is made;
*/

func InitCentroidLlyod(points []mapreduce.Points, numCentroid int) []mapreduce.Centroids {
	centroids := make([]mapreduce.Centroids, numCentroid)

	for i := 0; i < numCentroid; i++ {
		min := i * len(points) / numCentroid
		max := ((i + 1) * len(points)) / numCentroid
		centroids[i].Index = i
		centroids[i].Centroid = points[min:max][rand.Intn(len(points[min:max]))].Point
	}

	return centroids
}

func Llyod(numWorker int, numCentroid int, points []mapreduce.Points, algo string, clients []*rpc.Client, calls []*rpc.Call) {
	//creazione di punti e centroidi
	centroids := createInitValue(algo, numCentroid, points)

	log.Printf("Punti: %d", len(points))
	log.Printf("Centroidi: %d", len(centroids))

	jobMap := splitJobMap(points, numWorker)
	var changes []int
	var isChanged bool

	var it int
	log.Printf("Inizio iterazione dell'algoritmo")
	for {
		channel := make(chan []mapreduce.Clusters)
		clusterChannel := make([]mapreduce.Clusters, len(centroids))
		clusters := make([]mapreduce.Clusters, len(centroids))

		for i := 0; i < len(points); i++ {
			points[i].Centroids = centroids
		}

		//Assegnazione righe ai mapper.
		for i := range jobMap {
			go assignMap(jobMap[i], clients[i], channel)
			clusterChannel = <-channel

			for j := range clusterChannel {
				clusters[j].Centroid = clusterChannel[j].Centroid
				clusters[j].Changes += clusterChannel[j].Changes
				for k := range clusterChannel[j].PointsData {
					clusters[j].PointsData = append(clusters[j].PointsData, clusterChannel[j].PointsData[k])
				}
			}
		}

		//Se non si verificano cambiamenti nel clusters, l'algoritmo termina
		changes, isChanged = checkChanges(clusters, changes)
		if isChanged {
			log.Println("Numero di iterazioni totali: ", it)
			break
		}
		it++

		log.Printf("Iterazione numero: %d", it)
		for ii := range clusters {
			if len(clusters[ii].PointsData) != 0 {
				log.Printf("Cluster %d con %d punti", clusters[ii].Centroid.Index, len(clusters[ii].PointsData))
			}
		}

		jobReduce := splitJobReduce(clusters, numWorker)
		newCentroids := make([]mapreduce.Centroids, 0)
		centroids = nil

		for i := range jobReduce {
			calls[i] = clients[i].Go("API.Reduce", jobReduce[i], &newCentroids, nil)
			calls[i] = <-calls[i].Done
			for j := range newCentroids {
				centroids = append(centroids, newCentroids[j])
			}
		}
	}
}
