package main

import (
	"math"
)

// Mapper
/*Nella fase di Map, avviene la classificazione dei punti: ciascun Mapper in parallelo riceve in input un
sottoinsieme (o chunk) di punti dell’insieme di partenza e per ciascuno di questi punti calcola la distanza
euclidea tra il punto ed i k centroidi, identificando così il centroide che minimizza la distanza
e al cui cluster il punto viene assegnato.
*/

// Near Trova il punto più vicino tra punto e centroide e
//ritorna la distanza con il cluster appartenente
func Near(point Points, centroids []Centroids, distanceFunction DistanceMethod) (int, float64) {
	indexOfCluster := 0
	minSquaredDistance, _ := distanceFunction(point.Point, centroids[0].Centroid)
	for i := 1; i < len(centroids); i++ {
		squaredDistance, _ := distanceFunction(point.Point, centroids[i].Centroid)
		if squaredDistance < minSquaredDistance {
			minSquaredDistance = squaredDistance
			indexOfCluster = i
		}
	}
	return indexOfCluster, math.Sqrt(minSquaredDistance)
}

// Kmeans Algorithm Per ogni punto dell'insieme identifica il cluster più vicino
func Kmeans(points Points, distanceFunction DistanceMethod, cluster []Clusters) ([]Clusters, Points, error) {
	centroid := make([]Centroids, 0)
	for i := range cluster {
		centroid = append(centroid, cluster[i].Centroid)
	}
	closestCluster, _ := Near(points, centroid, distanceFunction)

	for i := range cluster {
		if closestCluster != points.ClusterNumber {
			cluster[i].Changes++
		}

		if i == closestCluster {
			points.ClusterNumber = closestCluster
			cluster[i].PointsData = append(cluster[i].PointsData, points)
		}
	}

	return cluster, points, nil
}
