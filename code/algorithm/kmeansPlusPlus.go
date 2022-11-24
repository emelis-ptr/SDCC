package algorithm

import (
	"SDCC-project/code/mapreduce"
	"SDCC-project/code/util"
	"log"
	"math/rand"
	"net/rpc"
	"sync"
)

/**
1. Initialize M as {µ1}, where µ1 is chosen uniformly at random from X.
2. Apply the MapReduce k-means++Map and k-means++Reduce to X.
3. Add the resulting point x to M.
4. Broadcast the new set M to each machine on the cluster.
5. Repeat steps 2 through 4 a total of k −1 times to produce k initial means.
6. Apply the standard k-means MapReduce algorithm, initialized with these means.
*/

var wg sync.WaitGroup

// KMeansPlusPlus : Algoritmo Kmeans
func KMeansPlusPlus(numWorker int, numCentroid int, numMapper int, numReducer int, points []mapreduce.Points, centroids []mapreduce.Centroids, clients []*rpc.Client, calls []*rpc.Call, testing bool) {
	log.Printf("Punti: %d", len(points))

	newCentroids := getCentroids(points, numWorker, numCentroid, centroids, clients, calls)
	log.Println("Fine iterazione dei centroidi ")

	//Stesso procedimento dell'algoritmo Llyod
	Llyod(numWorker, numMapper, numReducer, points, newCentroids, clients, calls, testing)
}

// InitCentroidKMeansPlusPlus : determina i centroidi iniziali
func InitCentroidKMeansPlusPlus(points []mapreduce.Points) []mapreduce.Centroids {
	centroids := make([]mapreduce.Centroids, 0)

	//Si sceglie casualmente il primo centroide
	cen := mapreduce.Centroids{Index: 0, Centroid: points[rand.Intn(len(points))].Point}
	centroids = append(centroids, cen)

	return centroids
}

// Iterazione per la scelta dei centroidi per l'algoritmo kmeans++
func getCentroids(points []mapreduce.Points, numWorker int, numCentroid int, centroids []mapreduce.Centroids, clients []*rpc.Client, calls []*rpc.Call) []mapreduce.Centroids {
	jobMap := util.SplitJobMap(points, numWorker)

	log.Println("Inizio iterazione dei centroidi")
	for {
		worker := rand.Intn(numWorker)
		for i := 0; i < len(points); i++ {
			points[i].Centroids = centroids
		}

		pointsChannel := make(chan []mapreduce.Points, len(points))
		pointsCh := make([]mapreduce.Points, len(points))
		for i := range jobMap {
			wg.Add(1)
			go assignMapKMeans(jobMap[i], clients[i], pointsChannel)
			pointsCh = <-pointsChannel
			for j := range pointsCh {
				points[j].Distance = pointsCh[j].Distance
			}
		}

		var newCentroids []mapreduce.Centroids

		calls[worker] = clients[worker].Go("API.ReduceKMeans", points, &newCentroids, nil)
		calls[worker] = <-calls[worker].Done

		centroids = append(centroids, newCentroids[0])

		if len(centroids) == numCentroid {
			break
		}
	}
	wg.Wait()
	return centroids
}

// Invia un insieme di punti al worker che esegue il job tramite la chiamata API
func assignMapKMeans(points []mapreduce.Points, client *rpc.Client, ch chan []mapreduce.Points) {
	defer wg.Done()
	var c []mapreduce.Points
	err := client.Call("API.MapperKMeans", points, &c)
	if err != nil {
		log.Fatal("Error in API.MapperKmeans: ", err)
	}
	ch <- c
}
