package mapreduce

import (
	"fmt"
	"log"
	"math"
)

//Points Point con il numero cluster di appartenza
type Points struct {
	ClusterNumber Centroids
	Centroids     []Centroids
	Point         []float64
	Distance      float64
}

type Clusters struct {
	Centroid   Centroids
	PointsData []Points
	Changes    int
}

type Centroids struct {
	Index    int
	Centroid []float64
}

// API service for RPC
type API int

func (a *API) Mapper(point []Points, cluster *[]Clusters) error {
	fmt.Println(" ")
	log.Printf(" ** Map phase **")
	log.Printf("Numero di punti assegnati %d", len(point))

	lenCentroid := len(point[0].Centroids)
	clusters := make([]Clusters, lenCentroid)

	for _, points := range point {
		closestCluster := 0
		var minSquaredDistance float64

		distance := 0.
		for ii := range points.Point {
			distance += (points.Point[ii] - points.Centroids[0].Centroid[ii]) * (points.Point[ii] - points.Centroids[0].Centroid[ii])
		}
		minSquaredDistance = math.Sqrt(distance)

		for i := 0; i < len(points.Centroids); i++ {
			distance1 := 0.
			for ii := range points.Point {
				distance1 += (points.Point[ii] - points.Centroids[i].Centroid[ii]) * (points.Point[ii] - points.Centroids[i].Centroid[ii])
			}
			squaredDistance := math.Sqrt(distance1)
			if squaredDistance < minSquaredDistance {
				minSquaredDistance = squaredDistance
				closestCluster = i
			}
		}

		for i := range clusters {
			clusters[i].Centroid.Index = i
			clusters[i].Centroid.Centroid = points.Centroids[i].Centroid

			if i == closestCluster {
				if closestCluster != points.ClusterNumber.Index {
					clusters[i].Changes++
				}
				points.ClusterNumber.Index = closestCluster
				clusters[i].PointsData = append(clusters[i].PointsData, points)
			}
		}
	}

	*cluster = clusters
	return nil
}

// MapperKMeans
//The Map phase operates on each point x in the dataset. For a given x,
//we compute the squared distance between x and each mean in M and find the
//minimum such squared distance D(x). We then emit a single value (x, D(x)),
//with no key/**
func (a *API) MapperKMeans(point []Points, cluster *[]Clusters) error {
	fmt.Println(" ")
	log.Printf(" ** Map phase **")
	log.Printf("Numero di punti assegnati %d", len(point))

	lenCentroid := len(point[0].Centroids)
	clusters := make([]Clusters, lenCentroid)

	for _, points := range point {
		closestCluster := 0
		var minSquaredDistance float64

		distance := 0.
		for ii := range points.Point {
			distance += (points.Point[ii] - points.Centroids[0].Centroid[ii]) * (points.Point[ii] - points.Centroids[0].Centroid[ii])
		}
		minSquaredDistance = math.Sqrt(distance)

		if len(points.Centroids) != 1 {
			for i := 0; i < len(points.Centroids); i++ {
				distance1 := 0.
				for ii := range points.Point {
					distance1 += (points.Point[ii] - points.Centroids[i].Centroid[ii]) * (points.Point[ii] - points.Centroids[i].Centroid[ii])
				}
				squaredDistance := math.Sqrt(distance1)
				if squaredDistance < minSquaredDistance {
					minSquaredDistance = squaredDistance
					closestCluster = i
				}
			}
		}
		points.Distance = minSquaredDistance

		for i := range clusters {
			clusters[i].Centroid.Index = i
			clusters[i].Centroid.Centroid = points.Centroids[i].Centroid

			if i == closestCluster {
				if closestCluster != points.ClusterNumber.Index {
					clusters[i].Changes++
				}
				points.ClusterNumber.Index = closestCluster
				clusters[i].PointsData = append(clusters[i].PointsData, points)
			}
		}
	}

	*cluster = clusters
	return nil
}
