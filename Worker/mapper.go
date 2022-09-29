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

// Mapper
/*Nella fase di Map, avviene la classificazione dei punti: ciascun Mapper in parallelo riceve in input un
sottoinsieme (o chunk) di punti dell’insieme di partenza e per ciascuno di questi punti calcola la distanza
euclidea tra il punto ed i k centroidi, identificando così il centroide che minimizza la distanza
e al cui cluster il punto viene assegnato.
*/
func (a *API) Mapper(clusteredPoint []ClusteredPoint, centroid *[]Point) error {

	_, err := Kmeans(clusteredPoint, *centroid, EuclideanDistance)

	return err
}

// Near Trova il punto più vicino tra punto e centroide e
//ritorna la distanza con il cluster appartenente
func Near(point ClusteredPoint, centroid []Point, distanceFunction DistanceMethod) (int, float64) {
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

// Kmeans Algorithm Per ogni punto dell'insieme identifica il cluster più vicino
func Kmeans(data []ClusteredPoint, centroid []Point, distanceFunction DistanceMethod) ([]ClusteredPoint, error) {

	for ii, jj := range data {
		closestCluster, _ := Near(jj, centroid, distanceFunction)
		data[ii].ClusterNumber = closestCluster
	}

	var changes int
	for ii, p := range data {
		if closestCluster, _ := Near(p, centroid, distanceFunction); closestCluster != p.ClusterNumber {
			changes++
			data[ii].ClusterNumber = closestCluster
		}
	}

	return data, nil
}
