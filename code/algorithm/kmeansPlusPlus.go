package algorithm

import (
	"SDCC-project/code/mapreduce"
	"SDCC-project/code/util"
	"log"
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

// KMeansPlusPlus : Algoritmo Kmeans
func KMeansPlusPlus(numWorker int, numCentroid int, numMapper int, points []mapreduce.Points, clients []*rpc.Client, calls []*rpc.Call, testing bool) {
	centroids := InitCentroidKMeansPlusPlus(points)

	log.Printf("Punti: %d", len(points))

	jobMap := util.SplitJobMap(points, numWorker)

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
			go assignJobsMapKMeans(jobMap[i], clients[i], channel, numMapper)
			clusterChannel = <-channel

			for j := range clusterChannel {
				clusters[j].Centroid = clusterChannel[j].Centroid
				clusters[j].Changes += clusterChannel[j].Changes
				for k := range clusterChannel[j].PointsData {
					clusters[j].PointsData = append(clusters[j].PointsData, clusterChannel[j].PointsData[k])
				}
			}
		}

		if len(centroids) == numCentroid {
			log.Println("Numero di iterazioni totali: ", it)
			break
		}

		jobReduce := util.SplitJobReduce(clusters, numWorker)
		newCentroids := make([]mapreduce.Centroids, 0)
		centroids = nil

		for i := range jobReduce {
			calls[i] = clients[i].Go("API.ReduceKMeans", jobReduce[i], &newCentroids, nil)
			calls[i] = <-calls[i].Done
			for j := range newCentroids {
				centroids = append(centroids, newCentroids[j])
			}
		}

		it++
		if testing == false {
			for ii := range clusters {
				log.Printf("Cluster %d con %d punti", clusters[ii].Centroid.Index, len(clusters[ii].PointsData))
			}
		}
	}
}

// InitCentroidKMeansPlusPlus : determina i centroidi iniziali
func InitCentroidKMeansPlusPlus(points []mapreduce.Points) []mapreduce.Centroids {
	centroids := make([]mapreduce.Centroids, 0)

	//Si sceglie casualmente il primo centroide
	cen := mapreduce.Centroids{Index: 0, Centroid: points[rand.Intn(len(points))].Point}
	centroids = append(centroids, cen)

	return centroids
}

// Suddivide i punti in base al numero di mapper che si vuole utilizzare
func assignJobsMapKMeans(points []mapreduce.Points, client *rpc.Client, ch chan []mapreduce.Clusters, numMapper int) {
	channel := make(chan []mapreduce.Clusters)
	clusterChannel := make([]mapreduce.Clusters, len(points[0].Centroids))
	clusters := make([]mapreduce.Clusters, len(points[0].Centroids))
	jobMap := util.SplitJobMap(points, numMapper)

	for i := range jobMap {
		go assignMapKMeans(jobMap[i], client, channel)
		clusterChannel = <-channel

		for j := range clusterChannel {
			clusters[j].Centroid = clusterChannel[j].Centroid
			clusters[j].Changes += clusterChannel[j].Changes
			for k := range clusterChannel[j].PointsData {
				clusters[j].PointsData = append(clusters[j].PointsData, clusterChannel[j].PointsData[k])
			}
		}
	}

	ch <- clusters
}

// Invia un insieme di punti al worker che esegue il job tramite la chiamata API
func assignMapKMeans(points []mapreduce.Points, client *rpc.Client, ch chan []mapreduce.Clusters) {
	var c []mapreduce.Clusters
	err := client.Call("API.MapperKMeans", points, &c)
	if err != nil {
		log.Fatal("Error in API.Mapper: ", err)
	}
	ch <- c
}
