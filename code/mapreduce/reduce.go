package mapreduce

import (
	"fmt"
	"log"
	"math/rand"
)

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
