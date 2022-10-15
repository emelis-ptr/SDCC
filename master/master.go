package main

import (
	"fmt"
	"log"
	"main/mapreduce"
	"main/util"
	"net/rpc"
	"os"
	"strconv"
)

type DistanceMethod func(firstVector, secondVector []float64) (float64, error)

func createInitValue() ([]mapreduce.Points, []mapreduce.Centroids) {
	numPoint, _ := strconv.Atoi(os.Getenv("NUMPOINT"))
	numCentroid, _ := strconv.Atoi(os.Getenv("NUMCENTROID"))
	numVector, _ := strconv.Atoi(os.Getenv("NUMVECTOR"))

	data := GeneratePoint(numPoint, numVector)
	points := CreateClusteredPoint(data)
	centroids := InitCentroid(points, numCentroid)

	return points, centroids
}

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

// Determina quanti cluster inviare al worker
func splitJobReduce(clusters []mapreduce.Clusters, numWorker int) [][]mapreduce.Clusters {
	var result [][]mapreduce.Clusters
	for i := 0; i < numWorker; i++ {

		min := i * len(clusters) / numWorker
		max := ((i + 1) * len(clusters)) / numWorker
		result = append(result, clusters[min:max])
	}
	return result
}

func main() {
	fmt.Println("Master is up")
	var err error
	var num int
	var conf util.Conf
	conf.ReadConf()

	util.OpenEnv()
	maxInteration, _ := strconv.Atoi(os.Getenv("MAXINTERAZIONI"))
	num, _ = strconv.Atoi(os.Getenv("NUM_WORKER"))

	//creazione di punti e centroidi
	points, centroids := createInitValue()

	clients := make([]*rpc.Client, num)
	call := make([]*rpc.Call, num)

	for i := 0; i < num; i++ {
		clients[i], err = rpc.DialHTTP("tcp", conf.PeerIP+":"+strconv.Itoa(conf.PeerPort))
		if err != nil {
			fmt.Println("Errore di connessione , retring ... ")
		}
	}

	fmt.Println(conf.PeerIP + " " + strconv.Itoa(conf.PeerPort))
	fmt.Println(clients)

	numWorker := len(clients)

	jobMap := splitJobMap(points, numWorker)
	var changes []int

	log.Printf("Inizio iterazione dell'algoritmo")
	for it := 0; it < maxInteration; it++ {
		c := make(chan []mapreduce.Clusters)
		clusterWorker := make([]mapreduce.Clusters, len(centroids))
		cluster := make([]mapreduce.Clusters, len(centroids))

		for i := 0; i < len(points); i++ {
			points[i].Centroids = centroids
		}

		//Assegnazione righe ai mapper.
		for i := range jobMap {
			go assignMap(jobMap[i], clients[i], c)
			clusterWorker = <-c
			for j := range clusterWorker {
				cluster[j].Centroid = clusterWorker[j].Centroid
				cluster[j].Changes += clusterWorker[j].Changes
				for k := range clusterWorker[j].PointsData {
					cluster[j].PointsData = append(cluster[j].PointsData, clusterWorker[j].PointsData[k])
				}
			}
		}

		jobReduce := splitJobReduce(cluster, numWorker)
		newCentroids := make([]mapreduce.Centroids, 0)
		centroids = nil

		for i := range jobReduce {
			call[i] = clients[i].Go("API.Reduce", jobReduce[i], &newCentroids, nil)
			call[i] = <-call[i].Done
			for j := range newCentroids {
				centroids = append(centroids, newCentroids[j])
			}
		}

		//Se non si verificano cambiamenti nel cluster, l'algoritmo termina
		changes = checkChanges(cluster, changes, numWorker, clients)

	}
}
