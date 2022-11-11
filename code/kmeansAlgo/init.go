package kmeansAlgo

import (
	"log"
	"main/code/mapreduce"
	"main/code/util"
	"net/rpc"
)

// Invia un insieme di punti al worker
func assignMap(points []mapreduce.Points, client *rpc.Client, ch chan []mapreduce.Clusters) {
	var c []mapreduce.Clusters
	err := client.Call("API.Mapper", points, &c)
	if err != nil {
		log.Fatal("Error in API.Mapper: ", err)
	}
	ch <- c
}

// Determina quanti punti inviare al worker
func splitJobMap(points []mapreduce.Points, numWorker int) [][]mapreduce.Points {
	var result [][]mapreduce.Points
	for i := 0; i < numWorker; i++ {

		min := i * len(points) / numWorker
		max := ((i + 1) * len(points)) / numWorker
		result = append(result, points[min:max])
	}
	return result
}

// Invia un insieme di punti al worker
func assignMapKMeans(points []mapreduce.Points, client *rpc.Client, ch chan []mapreduce.Clusters) {
	var c []mapreduce.Clusters
	err := client.Call("API.MapperKMeans", points, &c)
	if err != nil {
		log.Fatal("Error in API.Mapper: ", err)
	}
	ch <- c
}

func splitJobReduce(clusters []mapreduce.Clusters, numWorker int) [][]mapreduce.Clusters {
	var result [][]mapreduce.Clusters

	for i := 0; i < numWorker; i++ {
		min := i * len(clusters) / numWorker
		max := ((i + 1) * len(clusters)) / numWorker
		result = append(result, clusters[min:max])
	}

	return result
}

func createInitValue(algo string, numCentroid int, points []mapreduce.Points) []mapreduce.Centroids {
	var centroids []mapreduce.Centroids
	if algo == util.Llyod {
		centroids = InitCentroidLlyod(points, numCentroid)
	} else if algo == util.Standard {
		centroids = InitCentroidStandard(points, numCentroid)
	}

	return centroids
}

// checkChanges verifica se ci sono cambiamenti all'interno del cluster
func checkChanges(cluster []mapreduce.Clusters, changes []int) ([]int, bool) {
	isChanged := false
	var countChanges int
	for j := range cluster {
		for i, value := range changes {
			if i == j && value == cluster[j].Changes {
				countChanges++
			}
		}

		if countChanges == len(cluster) {
			isChanged = true
		}
	}

	changes = nil
	for i := range cluster {
		changes = append(changes, cluster[i].Changes)
	}

	return changes, isChanged
}
