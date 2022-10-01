package main

import (
	"math/rand"
)

//InitCentroid Si utilizza k-means++ per determinare i centroidi iniziali
func InitCentroid(data []Points, numCentroid int, distanceFunction DistanceMethod) []Centroid {
	centroids := make([]Centroid, numCentroid)
	centroidPoint := make([]Point, numCentroid)

	//Si sceglie casualmente il primo centroide
	centroidPoint[0] = data[rand.Intn(len(data))].Point
	centroids[0].Index = 0
	centroids[0].Centroid = centroidPoint[0]

	distance := make([]float64, len(data))

	wcss := make([]float64, 0)

	for ii := 1; ii < numCentroid; ii++ {
		var sum float64
		/* Per ogni punto trova il centroide più vicino, e salva la sua distanza in un array, ottendo la
		somma di tutte le distanze */
		for jj, p := range data {
			_, minDistance := Near(p, centroids[:ii], distanceFunction)
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

		centroidPoint[ii] = data[rand.Intn(len(data))].Point

		centroids[ii].Index = ii
		centroids[ii].Centroid = data[rand.Intn(len(data))].Point
	}

	lineChart(wcss)
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
