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

	time.Sleep(time.Second)
	err = client.Call("Registry.RetrieveMember", bools, &registrations)
	for err != nil {
		log.Fatal("Error call registry", err.Error())
	}

	for i := range registrations.Peer {
		log.Printf("Port %s", strconv.Itoa(registrations.Peer[i].Port))
		log.Printf("Address %s ", registrations.Peer[i].Address)
	}

	//creazione di punti e centroidi
	points, centroids := createInitValue()

	log.Printf("Punti: %d", len(points))
	log.Printf("Centroidi: %d", len(centroids))

	clients := make([]*rpc.Client, num)
	calls := make([]*rpc.Call, num)

	for i := 0; i < len(registrations.Peer); i++ {
		clients[i], err = rpc.DialHTTP("tcp", registrations.Peer[i].Address+":"+strconv.Itoa(registrations.Peer[i].Port))
		if err != nil {
			fmt.Println("Errore di connessione , retring ... ")
		}
	}

	numWorker := len(clients)

	jobMap := splitJobMap(points, numWorker)
	var changes []int
	var isChanged bool

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
			if len(cluster[ii].PointsData) != 0 {
				log.Printf("Cluster %d con %d punti", cluster[ii].Centroid.Index, len(cluster[ii].PointsData))
			}
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
		changes, isChanged = checkChanges(cluster, changes)

		if isChanged {
			writeFile(cluster)
			break
		}
	}
}

func writeFile(cluster []mapreduce.Clusters) {
	// create file
	f, err := os.Create("kmeans.txt")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file
	defer f.Close()

	for _, line := range cluster {
		_, err = f.WriteString(strconv.Itoa(line.Centroid.Index) + " - ")
		for i := range line.Centroid.Centroid {
			_, err = f.WriteString(strconv.FormatFloat(line.Centroid.Centroid[i], 'f', 5, 64) + " ")
		}
		_, err = f.WriteString("\n")
		for i := range line.PointsData {
			for j := range line.PointsData[i].Point {
				_, err = f.WriteString(strconv.FormatFloat(line.PointsData[i].Point[j], 'f', 5, 64) + " ")
			}
			_, err = f.WriteString(" - ")
		}
		_, err = f.WriteString("\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
