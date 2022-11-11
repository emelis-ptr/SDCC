package kmeansAlgo

import (
	"log"
	"main/code/mapreduce"
	"math/rand"
	"net/rpc"
)

/**
1. Initialize M as {µ1}, where µ1 is chosen uniformly at random from X.
2. Apply the MapReduce k-means++Map and k-means++Reduce to X.
3. Add the resulting point x to M.
4. Broadcast the new set M to each machine on the cluster.
5. Repeat steps 2 through 4 a total of k −1 times to produce k initial means.
6. Apply the standard k-means MapReduce algorithm, initialized with these means.
*/

func KMeansPlusPlus(numWorker int, numCentroid int, points []mapreduce.Points, algo string, clients []*rpc.Client, calls []*rpc.Call) {
	centroids := InitCentroidKMeansPlusPlus(points)

	log.Printf("Punti: %d", len(points))

	jobMap := splitJobMap(points, numWorker)

	var it int
	log.Printf("Inizio iterazione dell'algoritmo")
	for {
		log.Printf("Centroidi: %d", len(centroids))

		channel := make(chan []mapreduce.Clusters)
		clusterChannel := make([]mapreduce.Clusters, len(centroids))
		clusters := make([]mapreduce.Clusters, len(centroids))

		for i := 0; i < len(points); i++ {
			points[i].Centroids = centroids
		}

		//Assegnazione righe ai mapper.
		for i := range jobMap {
			go assignMapKMeans(jobMap[i], clients[i], channel)
			clusterChannel = <-channel

			for j := range clusterChannel {
				clusters[j].Centroid = clusterChannel[j].Centroid
				clusters[j].Changes += clusterChannel[j].Changes
				for k := range clusterChannel[j].PointsData {
					clusters[j].PointsData = append(clusters[j].PointsData, clusterChannel[j].PointsData[k])
				}
			}
		}

		jobReduce := splitJobReduce(clusters, numWorker)
		newCentroids := make([]mapreduce.Centroids, 0)
		centroids = nil

		for i := range jobReduce {
			calls[i] = clients[i].Go("API.ReduceKMeans", jobReduce[i], &newCentroids, nil)
			calls[i] = <-calls[i].Done
			for j := range newCentroids {
				centroids = append(centroids, newCentroids[j])
			}
		}

		if len(centroids) == numCentroid {
			log.Println("Numero di iterazioni totali: ", it)
			break
		}
		it++
		for ii := range clusters {
			if len(clusters[ii].PointsData) != 0 {
				log.Printf("Cluster %d con %d punti", clusters[ii].Centroid.Index, len(clusters[ii].PointsData))
			}
		}

	}
}

//InitCentroidKMeansPlusPlus Si utilizza k-means++ per determinare i centroidi iniziali
func InitCentroidKMeansPlusPlus(points []mapreduce.Points) []mapreduce.Centroids {
	centroids := make([]mapreduce.Centroids, 0)

	//Si sceglie casualmente il primo centroide
	cen := mapreduce.Centroids{Index: 0, Centroid: points[rand.Intn(len(points))].Point}
	centroids = append(centroids, cen)

	return centroids
}
