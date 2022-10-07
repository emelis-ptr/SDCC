package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math"
	"math/rand"
	"net/rpc"
	"os"
	"strconv"
)

// Points Point con il numero cluster di appartenza
type Points struct {
	ClusterNumber Centroids
	Centroids     []Centroids
	Point         []float64
}

type Clusters struct {
	Centroid   Centroids
	PointsData []Points
	Changes    int
}

type Centroids struct {
	Index    int
	Centroid []float64
}

type DistanceMethod func(firstVector, secondVector []float64) (float64, error)

//InitCentroid Si utilizza k-means++ per determinare i centroidi iniziali
func InitCentroid(points []Points, numCentroid int) []Centroids {
	centroids := make([]Centroids, numCentroid)

	//Si sceglie casualmente il primo centroide
	centroids[0].Index = 0
	centroids[0].Centroid = points[rand.Intn(len(points))].Point

	distance := make([]float64, len(points))

	wcss := make([]float64, 0)

	for ii := 1; ii < numCentroid; ii++ {
		var sum float64
		// Per ogni punto trova il centroide più vicino, e salva la sua distanza in un array, ottendo la
		//somma di tutte le distanze
		for jj, p := range points {
			_, minDistance := Near(p, centroids[:ii], EuclideanDistance)
			distance[jj] = minDistance * minDistance
			sum += distance[jj]
		}
		wcss = append(wcss, sum)
		//Trova una distanza casuale moltiplicando un valore random con la somma delle distanze
		randomDistance := rand.Float64() * sum
		jj := 0

		// Assegna i centroidi.
		for sum = distance[0]; sum < randomDistance; sum += distance[jj] {
			jj++
		}

		centroids[ii].Index = ii
		centroids[ii].Centroid = points[rand.Intn(len(points))].Point
	}

	return centroids
}

// GeneratePoint Genera un insieme di punti con una dimensione data
func GeneratePoint(numPoint int, numVector int) [][]float64 {
	data := make([][]float64, numPoint)
	for i := 0; i < numPoint; i++ {
		for j := 0; j < numVector; j++ {
			data[i] = append(data[i], rand.Float64())
		}
	}
	return data
}

func CreateClusteredPoint(rawData [][]float64) []Points {
	data := make([]Points, len(rawData))
	for ii, jj := range rawData {
		data[ii].Point = jj
	}
	return data
}

// Near Trova il punto più vicino tra punto e centroide e
//ritorna la distanza con il cluster appartenente
func Near(point Points, cluster []Centroids, distanceFunction DistanceMethod) (int, float64) {
	indexOfCluster := 0
	minSquaredDistance, _ := distanceFunction(point.Point, cluster[0].Centroid)
	for i := 1; i < len(cluster); i++ {
		squaredDistance, _ := distanceFunction(point.Point, cluster[i].Centroid)
		if squaredDistance < minSquaredDistance {
			minSquaredDistance = squaredDistance
			indexOfCluster = i
		}
	}
	return indexOfCluster, math.Sqrt(minSquaredDistance)
}

//EuclideanDistance Determina la distanza euclidea tra due punti
func EuclideanDistance(firstVector, secondVector []float64) (float64, error) {
	distance := 0.
	for ii := range firstVector {
		distance += (firstVector[ii] - secondVector[ii]) * (firstVector[ii] - secondVector[ii])
	}
	return math.Sqrt(distance), nil
}

// Invia un insieme di punti al worker
func assignMap(points []Points, client *rpc.Client, ch chan []Clusters) {
	var c []Clusters
	err := client.Call("API.Mapper", points, &c)
	if err != nil {
		log.Fatal("Error in API.Mapper: ", err)
	}
	ch <- c
}

// Determina quanti punti inviare al worker
func splitJobMap(points []Points, numWorker int) [][]Points {
	var result [][]Points
	for i := 0; i < numWorker; i++ {

		min := i * len(points) / numWorker
		max := ((i + 1) * len(points)) / numWorker
		result = append(result, points[min:max])
	}
	return result
}

// Determina quanti cluster inviare al worker
func splitJobReduce(clusters []Clusters, numWorker int) [][]Clusters {
	var result [][]Clusters
	for i := 0; i < numWorker; i++ {

		min := i * len(clusters) / numWorker
		max := ((i + 1) * len(clusters)) / numWorker
		result = append(result, clusters[min:max])
	}
	return result
}

func main() {
	var err error
	var num int
	var port int

	num, _ = strconv.Atoi(os.Args[1])
	clients := make([]*rpc.Client, num)
	call := make([]*rpc.Call, num)

	for i := 0; i < num; i++ {
		port = 4041 + i
		clients[i], err = rpc.DialHTTP("tcp", "localhost:"+strconv.Itoa(port))
		if err != nil {
			fmt.Println("Errore di connessione , retring ... ")
		}
	}

	//Apertura file.
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("failed to open")

	}
	err = godotenv.Load(file.Name())
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	numWorker := len(clients)
	numPoint, _ := strconv.Atoi(os.Getenv("NUMPOINT"))
	numCentroid, _ := strconv.Atoi(os.Getenv("NUMCENTROID"))
	numVector, _ := strconv.Atoi(os.Getenv("NUMVECTOR"))
	maxInteration, _ := strconv.Atoi(os.Getenv("MAXINTERAZIONI"))

	data := GeneratePoint(numPoint, numVector)
	points := CreateClusteredPoint(data)
	centroids := InitCentroid(points, numCentroid)

	jobMap := splitJobMap(points, numWorker)
	var changes []int

	for it := 0; it < maxInteration; it++ {
		c := make(chan []Clusters)
		clusterWorker := make([]Clusters, len(centroids))
		cluster := make([]Clusters, len(centroids))

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
		newCentroids := make([]Centroids, 0)
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

// checkChanges verifica se ci sono cambiamenti all'interno del cluster
func checkChanges(cluster []Clusters, changes []int, numWorker int, clients []*rpc.Client) []int {
	var countChanges int
	for j := range cluster {
		for _, value := range changes {
			if value == cluster[j].Changes {
				countChanges++
			}
		}

		if countChanges == len(cluster) {
			for i := 0; i < numWorker; i++ {
				err := clients[i].Close()
				if err != nil {
					return nil
				}
			}
		}
	}
	changes = nil
	for i := range cluster {
		changes = append(changes, cluster[i].Changes)
	}
	return changes
}
