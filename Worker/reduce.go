package main

import (
	"errors"
	"os"
)

//Reduce
/**
Nella fase di Reduce, avviene il calcolo dei nuovi centroidi: ciascun reducer in parallelo riceve in input
tutti i punti assegnati ad un determinato cluster e calcola il valore del centroide di quel cluster.
*/

// NewCentroid Determina i nuovi centroidi in base all'insieme dei punti del cluster
func NewCentroid(clusters []Clusters) []Centroids {
	clusters, err := deleteEmptyCluster(clusters)
	if err != nil {
		return nil
	}

	centroid := make([]Centroids, len(clusters))
	var lenPoint int

	for ii := range clusters {
		lenPoint = len(clusters[ii].PointsData[0].Point)

		p := make([][]float64, lenPoint)

		for j := range clusters[ii].PointsData {
			for k := range clusters[ii].PointsData[j].Point {
				p[k] = append(p[k], clusters[ii].PointsData[j].Point[k])
			}
		}

		var mean Point
		for k := range p {
			var sum float64
			for j := range p[k] {
				sum += p[k][j]
			}
			var op = sum / float64(len(p[k]))
			mean = append(mean, op)
		}

		centroid[ii].Index = ii
		centroid[ii].Centroid = mean
	}

	return centroid
}

// noChangeStop se per tutti i cluster non ci sono stati cambiamenti, l'algoritmo termina
func noChangesStop(clusters []Clusters) {
	var countChanges int
	for i := range clusters {
		if clusters[i].Changes == 0 {
			countChanges++
		}
		if countChanges == len(clusters) {
			os.Exit(0)
		}
	}
}

// deleteEmptyCluster Se esistiono cluster con un insieme di punti vuoto, viene eliminato
func deleteEmptyCluster(clusters []Clusters) ([]Clusters, error) {
	var index int
	for ii := range clusters {
		if len(clusters[ii].PointsData) == 0 {
			index = ii

			if index < 0 || index >= len(clusters) {
				return nil, errors.New("index cannot be less than 0")
			}

			clusters = append(clusters[:index], clusters[index+1:]...)
		}
	}

	return clusters, nil
}
