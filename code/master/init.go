package main

import (
	"main/code/mapreduce"
	"main/code/util"
	"math"
	"math/rand"
)

type DistanceMethod func(firstVector, secondVector []float64) (float64, error)

// GeneratePoint Genera un insieme di punti con una dimensione data
func GeneratePoint(numPoint int) [][]float64 {
	numVector := util.NumVector
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
