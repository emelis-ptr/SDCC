package util

import (
	"SDCC-project/code/mapreduce"
	"math/rand"
)

// GeneratePoint : Genera un insieme di punti
func GeneratePoint(numPoint int) [][]float64 {
	numVector := NumVector
	data := make([][]float64, numPoint)
	for i := 0; i < numPoint; i++ {
		for j := 0; j < numVector; j++ {
			data[i] = append(data[i], rand.Float64())
		}
	}
	return data
}

// CreateClusteredPoint : assegna al field Point il punto generato
func CreateClusteredPoint(rawData [][]float64) []mapreduce.Points {
	data := make([]mapreduce.Points, len(rawData))
	for ii, jj := range rawData {
		data[ii].Point = jj
	}
	return data
}

// SplitJobMap : Determina quanti punti inviare al worker
func SplitJobMap(points []mapreduce.Points, numWorker int) [][]mapreduce.Points {
	var result [][]mapreduce.Points
	for i := 0; i < numWorker; i++ {

		min := i * len(points) / numWorker
		max := ((i + 1) * len(points)) / numWorker
		result = append(result, points[min:max])
	}
	return result
}

// SplitJobReduce : Determina quanti punti inviare al worker
func SplitJobReduce(clusters []mapreduce.Clusters, numWorker int) [][]mapreduce.Clusters {
	var result [][]mapreduce.Clusters

	for i := 0; i < numWorker; i++ {
		min := i * len(clusters) / numWorker
		max := ((i + 1) * len(clusters)) / numWorker
		result = append(result, clusters[min:max])
	}

	return result
}

// CheckChanges verifica se si sono verificati cambiamenti all'interno del cluster
func CheckChanges(cluster []mapreduce.Clusters, changes []int) ([]int, bool) {
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
