package mapreduce

import (
	"fmt"
	"log"
	"math"
	"math/rand"
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

func (a *API) Reduce(clusters []Clusters, centroid *[]Centroids) error {
	fmt.Println(" ")
	log.Printf(" ** Reduce phase **")
	log.Printf("Numero di cluster assegnati %d", len(clusters))

	for ii := range clusters {
		log.Printf("Cluster %d con %d punti", clusters[ii].Centroid.Index, len(clusters[ii].PointsData))
	}

	//Crea centroid
	centroids := make([]Centroids, len(clusters))
	var lenPoint int

	if len(clusters) != 0 {
		for ii := range clusters {
			if len(clusters[ii].PointsData) != 0 {

				lenPoint = len(clusters[ii].PointsData[0].Point)

				p := make([][]float64, lenPoint)

				for j := range clusters[ii].PointsData {
					for k := range clusters[ii].PointsData[j].Point {
						p[k] = append(p[k], clusters[ii].PointsData[j].Point[k])
					}
				}

				var mean []float64
				for k := range p {
					var sum float64
					for j := range p[k] {
						sum += p[k][j]
					}
					var op = sum / float64(len(p[k]))
					mean = append(mean, op)
				}

				centroids[ii].Index = ii
				centroids[ii].Centroid = mean
			}
		}
	}
	*centroid = centroids

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

// ReduceKMeans
//We reduce the first element of two value pairs
//by choosing one of these elements with probability proportional to the second
//element in each pair, and reduce the second element of the pairs by summation.
func (a *API) ReduceKMeans(clusters []Clusters, centroid *[]Centroids) error {
	fmt.Println(" ")
	log.Printf(" ** Reduce phase **")
	log.Printf("Numero di cluster assegnati %d", len(clusters))

	for ii := range clusters {
		log.Printf("Cluster %d con %d punti", clusters[ii].Centroid.Index, len(clusters[ii].PointsData))
	}

	//Crea centroid
	centroids := make([]Centroids, len(clusters))
	var lenPoint int

	if len(clusters) != 0 {
		for ii := range clusters {
			if len(clusters[ii].PointsData) != 0 {

				lenPoint = len(clusters[ii].PointsData[0].Point)

				p := make([][]float64, lenPoint)

				for j := range clusters[ii].PointsData {
					for k := range clusters[ii].PointsData[j].Point {
						p[k] = append(p[k], clusters[ii].PointsData[j].Point[k])
					}
				}

				var mean []float64
				for k := range p {
					var sum float64
					for j := range p[k] {
						sum += p[k][j]
					}
					var op = sum / float64(len(p[k]))
					mean = append(mean, op)
				}

				centroids[ii].Index = ii
				centroids[ii].Centroid = mean
			}
		}
	}

	point := make([]Points, 0)
	if len(clusters) != 0 {
		for ii := range clusters {
			if len(clusters[ii].PointsData) != 0 {
				for i := range clusters[ii].PointsData {
					point = append(point, clusters[ii].PointsData[i])
				}
			}
		}
	}

	if (len(point)) != 0 {
		distance := make([]float64, len(point))
		var c Centroids
		var sum float64
		jj := 0

		for pp := range point {
			distance[pp] = point[pp].Distance * point[pp].Distance
			sum += distance[pp]
		}

		//Trova una distanza casuale moltiplicando un valore random con la somma delle distanze
		randomDistance := rand.Float64() * sum

		// Assegna i centroidi.
		for sum = distance[0]; sum < randomDistance; sum += distance[jj] {
			jj++
		}

		c.Index = len(clusters) + 1
		c.Centroid = point[jj].Point
		centroids = append(centroids, c)
	}

	*centroid = centroids

	return nil
}
