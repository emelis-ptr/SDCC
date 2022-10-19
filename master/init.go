package main

import (
	"main/mapreduce"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func createInitValue() ([]mapreduce.Points, []mapreduce.Centroids) {
	numPoint, _ := strconv.Atoi(os.Getenv("NUMPOINT"))
	numCentroid, _ := strconv.Atoi(os.Getenv("NUMCENTROID"))
	numVector, _ := strconv.Atoi(os.Getenv("NUMVECTOR"))

	data := GeneratePoint(numPoint, numVector)
	points := CreateClusteredPoint(data)
	centroids := InitCentroid(points, numCentroid)

	return points, centroids
}

//InitCentroid Si utilizza k-means++ per determinare i centroidi iniziali
func InitCentroid(points []mapreduce.Points, numCentroid int) []mapreduce.Centroids {
	centroids := make([]mapreduce.Centroids, numCentroid)

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

func CreateClusteredPoint(rawData [][]float64) []mapreduce.Points {
	data := make([]mapreduce.Points, len(rawData))
	for ii, jj := range rawData {
		data[ii].Point = jj
	}
	return data
}

// Near Trova il punto più vicino tra punto e centroide e
//ritorna la distanza con il cluster appartenente
func Near(point mapreduce.Points, cluster []mapreduce.Centroids, distanceFunction DistanceMethod) (int, float64) {
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

//verifica se ci sono cluster con un insieme di punti vuoto
func checkEmptyPoint(cluster []mapreduce.Clusters) {
	//Verica che non ci siano cluster vuoti, altrimenti lo elimina
	var index []int
	lenCluster := len(cluster)
	for ii := 0; ii < lenCluster; ii++ {
		if len(cluster[ii].PointsData) == 0 {
			index = append(index, cluster[ii].Centroid.Index)
		}
	}

}
