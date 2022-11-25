package algorithm

import (
	"SDCC-project/code/mapreduce"
	"SDCC-project/code/util"
	"log"
	"math/rand"
	"net/rpc"
	"time"
)

func CreateInitValue(algo string, numCentroid int, points []mapreduce.Points) []mapreduce.Centroids {
	var centroids []mapreduce.Centroids
	if algo == util.Llyod {
		centroids = InitCentroidLlyod(points, numCentroid)
	} else if algo == util.Standard {
		centroids = InitCentroidStandard(points, numCentroid)
	}

	return centroids
}

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

// InitCentroidLlyod : assegna i punti in cluster diversi
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

// Llyod : algoritmo
func Llyod(numWorker int, numMapper int, numReducer int, points []mapreduce.Points, centroids []mapreduce.Centroids, clients []*rpc.Client, calls []*rpc.Call, testing bool, algo string) {
	log.Printf("Worker: %d", len(clients))
	log.Printf("Punti: %d - Centroidi: %d", len(points), len(centroids))
	log.Printf("Mapper: %d - Reducer: %d", numMapper, numReducer)

	jobMap := util.SplitJobMap(points, numWorker)
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
			go assignJobsMap(jobMap[i], clients[i], channel, numMapper)
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
		changes, isChanged = util.CheckChanges(clusters, changes)
		if isChanged {
			log.Println("Numero di iterazioni totali: ", it)
			util.Plot(clusters, numMapper, numReducer, len(points))
			util.WriteClusters(clusters, len(points), numMapper, numReducer, algo)
			break
		}

		it++
		if !testing {
			log.Printf("Iterazione numero: %d", it)
			for ii := range clusters {
				log.Printf("Cluster %d con %d punti", clusters[ii].Centroid.Index, len(clusters[ii].PointsData))
			}
		}
		jobReduce := util.SplitJobReduce(clusters, numWorker)
		newCentroids := make([]mapreduce.Centroids, 0)
		channelR := make(chan []mapreduce.Centroids)
		centroids = nil

		for i := range jobReduce {
			jobReduce2 := util.SplitJobReduce(jobReduce[i], numReducer)

			for j := range jobReduce2 {
				go assignReduce(calls, clients, i, jobReduce2[j], channelR)
				newCentroids = <-channelR

				for k := range newCentroids {
					centroids = append(centroids, newCentroids[k])
				}
			}
		}
	}
}

// Suddivide i punti per ogni worker in base al numero di mapper specificato che si vuole utilizzare
func assignJobsMap(points []mapreduce.Points, client *rpc.Client, ch chan []mapreduce.Clusters, numMapper int) {
	channel := make(chan []mapreduce.Clusters)
	clusterChannel := make([]mapreduce.Clusters, len(points[0].Centroids))
	clusters := make([]mapreduce.Clusters, len(points[0].Centroids))
	jobMap := util.SplitJobMap(points, numMapper)

	for i := range jobMap {
		go assignMap(i, client, jobMap, channel)
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

// Ogni worker esegue una chiama API.Mapper
func assignMap(i int, client *rpc.Client, jobMap [][]mapreduce.Points, ch chan []mapreduce.Clusters) {
	time.Sleep(time.Second)
	var c []mapreduce.Clusters
	err := client.Call("API.Mapper", jobMap[i], &c)
	if err != nil {
		log.Fatal("Error in API.Mapper: ", err)
	}

	ch <- c
}

// Ogni worker esegue la chiamata API.Reduce
func assignReduce(calls []*rpc.Call, clients []*rpc.Client, i int, jobReduce []mapreduce.Clusters, ch chan []mapreduce.Centroids) {
	newCentroids := make([]mapreduce.Centroids, 0)

	calls[i] = clients[i].Go("API.Reduce", jobReduce, &newCentroids, nil)
	calls[i] = <-calls[i].Done

	ch <- newCentroids
}
