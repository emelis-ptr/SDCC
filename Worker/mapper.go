package main

import (
	"log"
	"math"
)

// Mapper
/*Nella fase di Map, avviene la classificazione dei punti: ciascun Mapper in parallelo riceve in input un
sottoinsieme (o chunk) di punti dell’insieme di partenza e per ciascuno di questi punti calcola la distanza
euclidea tra il punto ed i k centroidi, identificando così il centroide che minimizza la distanza
e al cui cluster il punto viene assegnato.
*/
func Mapper(centroid []Centroid, clusteredPoint *[]Points) ([]Cluster, error) {

	cp, err := Kmeans(*clusteredPoint, centroid, EuclideanDistance)
	if err != nil {
		log.Fatal(err)
	}

	return cp, nil
}

// Near Trova il punto più vicino tra punto e centroide e
//ritorna la distanza con il cluster appartenente
func Near(point Points, centroids []Centroid, distanceFunction DistanceMethod) (int, float64) {
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
func Kmeans(points []Points, centroids []Centroid, distanceFunction DistanceMethod) ([]Cluster, error) {
	cluster := make([]Cluster, len(centroids))

	for ii, jj := range points {
		closestCluster, _ := Near(jj, centroids, distanceFunction)
		for i := range centroids {
			if i == closestCluster {
				cluster[i].Centroid.Index = closestCluster
				cluster[i].Centroid.Centroid = centroids[closestCluster].Centroid
				cluster[i].PointsData = append(cluster[i].PointsData, points[ii])
			}
		}
	}

	return cluster, nil
}
