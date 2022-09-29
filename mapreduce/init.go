package main

import "math/rand"

//initCentroid Si utilizza k-means++ per determinare i centroidi iniziali
func initCentroid(data []ClusteredPoint, numCentroid int, distanceFunction DistanceMethod) []Point {
	centroid := make([]Point, numCentroid)

	//Si sceglie casualmente il primo centroide
	centroid[0] = data[rand.Intn(len(data))].Point

	distance := make([]float64, len(data))

	for ii := 1; ii < numCentroid; ii++ {
		var sum float64
		/* Per ogni punto trova il centroide più vicino, e salva la sua distanza in un array, ottendo la
		somma di tutte le distanze */
		for jj, p := range data {
			_, minDistance := near(p, centroid[:ii], distanceFunction)
			distance[jj] = minDistance
			sum += distance[jj]
		}
		//Trova una distanza casuale moltiplicando un valore random con la somma delle distanze
		randomDistance := rand.Float64() * sum
		jj := 0

		// Assegna i centroidi.
		for sum = distance[0]; sum < randomDistance; sum += distance[jj] {
			jj++
		}
		centroid[ii] = data[jj].Point
	}
	return centroid
}

// generatePoint Genera un insieme di punti con una dimensione data
func generatePoint(numPoint int, numVector int) [][]float64 {
	data := make([][]float64, numPoint)
	for i := 0; i < numPoint; i++ {
		for j := 0; j < numVector; j++ {
			data[i] = append(data[i], rand.Float64())
		}
	}
	return data
}

// createClusteredPoint
func createClusteredPoint(rawData [][]float64) []ClusteredPoint {
	data := make([]ClusteredPoint, len(rawData))
	for ii, jj := range rawData {
		data[ii].Point = jj
	}
	return data
}
