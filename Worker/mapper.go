package main

import (
	"log"
	"math"
)

// Points Point con il numero cluster di appartenza
type Points struct {
	ClusterNumber int
	Point         Point
}

type Cluster struct {
	Centroid   Centroid
	PointsData []Points
	Changes    int
}

type Centroid struct {
	Index    int
	Centroid Point
}

type DistanceMethod func(firstVector, secondVector []float64) (float64, error)

// Mapper
/*Nella fase di Map, avviene la classificazione dei punti: ciascun Mapper in parallelo riceve in input un
sottoinsieme (o chunk) di punti dell’insieme di partenza e per ciascuno di questi punti calcola la distanza
euclidea tra il punto ed i k centroidi, identificando così il centroide che minimizza la distanza
e al cui cluster il punto viene assegnato.
*/
func Mapper(points *Points, centroid []Centroid, cluster *[]Cluster) ([]Cluster, Points, error) {

	cp, p, err := Kmeans(*points, centroid, EuclideanDistance, *cluster)
	if err != nil {
		log.Fatal(err)
	}

	*cluster = cp
	*points = p
	return cp, p, nil
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
func Kmeans(points Points, centroids []Centroid, distanceFunction DistanceMethod, cluster []Cluster) ([]Cluster, Points, error) {

	closestCluster, _ := Near(points, centroids, distanceFunction)

	for i := range centroids {
		if closestCluster != points.ClusterNumber {
			cluster[i].Changes++
		}

		if i == closestCluster {
			points.ClusterNumber = closestCluster
			cluster[i].Centroid.Index = closestCluster
			cluster[i].Centroid.Centroid = centroids[closestCluster].Centroid
			cluster[i].PointsData = append(cluster[i].PointsData, points)
		}
	}

	return cluster, points, nil
}
