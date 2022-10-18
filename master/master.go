package main

import (
	"fmt"
	"log"
	"main/mapreduce"
	"main/util"
	"net/rpc"
	"os"
	"strconv"
	"time"
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

	var registrations util.Registration

	var bools bool
	client, err := rpc.DialHTTP("tcp", conf.RegIP+":"+strconv.Itoa(conf.RegPort))

	var try int
	for err != nil && try < 5 {
		//if the port is closed on first try, try again. Ten tries are allowed
		client, err = rpc.DialHTTP("tcp", conf.RegIP+":"+strconv.Itoa(conf.RegPort))
		try++
	}
	time.Sleep(time.Second)
	err = client.Call("Registry.RetrieveMember", bools, &registrations)

	for i := range registrations.Peer {
		log.Printf("Port %s", strconv.Itoa(registrations.Peer[i].Port))
		log.Printf("Address %s ", registrations.Peer[i].Address)
	}

	//creazione di punti e centroidi
	points, centroids := createInitValue()

	clients := make([]*rpc.Client, num)
	calls := make([]*rpc.Call, num)

	for i := 0; i < len(registrations.Peer); i++ {
		clients[i], err = rpc.DialHTTP("tcp", registrations.Peer[i].Address+":"+strconv.Itoa(registrations.Peer[i].Port))
		if err != nil {
			fmt.Println("Errore di connessione , retring ... ")
		}
	}

	if len(registrations.Peer) < num {
		time.Sleep(time.Second)
		return
	}
	numWorker := len(clients)

	jobMap := splitJobMap(points, numWorker)
	var changes []int

	log.Printf("Inizio iterazione dell'algoritmo")
	for it := 0; it < maxInteration; it++ {
		log.Printf("Iterazione numero: %d", it)

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

		for ii := range cluster {
			log.Printf("Cluster %d con %d punti", cluster[ii].Centroid.Index, len(cluster[ii].PointsData))
		}

		jobReduce := splitJobReduce(cluster, numWorker)
		newCentroids := make([]mapreduce.Centroids, 0)
		centroids = nil

		for i := range jobReduce {
			calls[i] = clients[i].Go("API.Reduce", jobReduce[i], &newCentroids, nil)
			calls[i] = <-calls[i].Done
			for j := range newCentroids {
				centroids = append(centroids, newCentroids[j])
			}
		}

		//Se non si verificano cambiamenti nel cluster, l'algoritmo termina
		changes = checkChanges(cluster, changes, numWorker, clients, client)
	}

}
