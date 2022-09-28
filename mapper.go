package main

import (
	"math"
)

// ClusteredPoint Point con il numero cluster di appartenza
type ClusteredPoint struct {
	ClusterNumber int
	Point
}

// DistanceMethod Calcola la distanza tra due punti
type DistanceMethod func(first, second []float64) (float64, error)

// mapper
/*Nella fase di Map, avviene la classificazione dei punti: ciascun mapper in parallelo riceve in input un
sottoinsieme (o chunk) di punti dell’insieme di partenza e per ciascuno di questi punti calcola la distanza
euclidea tra il punto ed i k centroidi, identificando così il centroide che minimizza la distanza
e al cui cluster il punto viene assegnato.
*/
func mapper(clusteredPoint []ClusteredPoint, centroid []Point, distanceFunction DistanceMethod) ([]ClusteredPoint, error) {

	clusteredData, err := kmeans(clusteredPoint, centroid, distanceFunction)

	return clusteredData, err
}

// Trova il punto più vicino tra punto e centroide e
//ritorna la distanza con il cluster appartenente
func near(point ClusteredPoint, centroid []Point, distanceFunction DistanceMethod) (int, float64) {
	indexOfCluster := 0
	minSquaredDistance, _ := distanceFunction(point.Point, centroid[0])
	for i := 1; i < len(centroid); i++ {
		squaredDistance, _ := distanceFunction(point.Point, centroid[i])
		if squaredDistance < minSquaredDistance {
			minSquaredDistance = squaredDistance
			indexOfCluster = i
		}
	}
	return indexOfCluster, math.Sqrt(minSquaredDistance)
}

// K-Means Algorithm Per ogni punto dell'insieme identifica il cluster più vicino
func kmeans(data []ClusteredPoint, centroid []Point, distanceFunction DistanceMethod) ([]ClusteredPoint, error) {

	for ii, jj := range data {
		closestCluster, _ := near(jj, centroid, distanceFunction)
		data[ii].ClusterNumber = closestCluster
	}

	var changes int
	for ii, p := range data {
		if closestCluster, _ := near(p, centroid, distanceFunction); closestCluster != p.ClusterNumber {
			changes++
			data[ii].ClusterNumber = closestCluster
		}
	}

	return data, nil
}
