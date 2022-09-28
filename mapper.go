package main

import (
	"math"
	"math/rand"
)

/**
Nella fase di Map, avviene la classificazione dei punti: ciascun mapper in parallelo riceve in input un
sottoinsieme (o chunk) di punti dell’insieme di partenza e per ciascuno di questi punti calcola la distanza
euclidea tra il punto ed i k centroidi, identificando così il centroide che minimizza la distanza
e al cui cluster il punto viene assegnato.
*/

// ClusteredPoint Abstracts the Point with a cluster number
// Update and computeation becomes more efficient
type ClusteredPoint struct {
	ClusterNumber int
	Point
}

// DistanceFunction To compute the distance between observations
type DistanceFunction func(first, second []float64) (float64, error)

// Find the closest observation and return the distance
// Index of observation, distance
func near(p ClusteredPoint, mean []Point, distanceFunction DistanceFunction) (int, float64) {
	indexOfCluster := 0
	minSquaredDistance, _ := distanceFunction(p.Point, mean[0])
	for i := 1; i < len(mean); i++ {
		squaredDistance, _ := distanceFunction(p.Point, mean[i])
		if squaredDistance < minSquaredDistance {
			minSquaredDistance = squaredDistance
			indexOfCluster = i
		}
	}
	return indexOfCluster, math.Sqrt(minSquaredDistance)
}

// Instead of initializing randomly the seeds, make a sound decision of initializing
func initCentroid(data []ClusteredPoint, numCentroid int, numVector int, distanceFunction DistanceFunction) []Point {
	centroid := make([]Point, numCentroid)
	centroid[0] = data[rand.Intn(len(data))].Point

	for ii := 1; ii < numCentroid; ii++ {
		for j := 0; j < numVector; j++ {
			centroid[ii] = append(centroid[ii], rand.Float64())
		}
	}
	/*d2 := make([]float64, len(data))
	for ii := 1; ii < numCentroid; ii++ {
		var sum float64
		for jj, p := range data {
			_, minDistance := near(p, centroid[:ii], distanceFunction)
			d2[jj] = minDistance * minDistance
			sum += d2[jj]
		}
		target := rand.Float64() * sum
		jj := 0
		for sum = d2[0]; sum < target; sum += d2[jj] {
			jj++
		}
		centroid[ii] = data[jj].Point
	}*/
	return centroid
}

// K-Means Algorithm
func kmeans(data []ClusteredPoint, centroid []Point, distanceFunction DistanceFunction, threshold int) ([]ClusteredPoint, error) {
	counter := 0
	print(counter, "\n")
	for ii, jj := range data {
		closestCluster, _ := near(jj, centroid, distanceFunction)
		data[ii].ClusterNumber = closestCluster
	}
	/*mLen := make([]int, len(centroid))
	for n := len(data[0].Point); ; {
		for ii := range centroid {
			centroid[ii] = make(Point, n)
			mLen[ii] = 0
		}
		for _, p := range data {
			centroid[p.ClusterNumber].Add(p.Point)
			mLen[p.ClusterNumber]++
		}
		for ii := range centroid {
			centroid[ii].Mul(1 / float64(mLen[ii]))
		}
		var changes int
		for ii, p := range data {
			if closestCluster, _ := near(p, centroid, distanceFunction); closestCluster != p.ClusterNumber {
				changes++
				data[ii].ClusterNumber = closestCluster
			}
		}
		counter++
		if changes == 0 || counter > threshold {
			return data, nil
		}
	}*/
	return data, nil
}

// mapper K-Means Algorithm with smart seeds
// as known as K-Means ++
func mapper(rawData [][]float64, numCentroid int, numVector int, distanceFunction DistanceFunction, threshold int) ([]ClusteredPoint, error) {
	data := make([]ClusteredPoint, len(rawData))
	for ii, jj := range rawData {
		data[ii].Point = jj
	}
	centroid := initCentroid(data, numCentroid, numVector, distanceFunction)
	for i := range centroid {
		print(centroid[i], " centroid \n")
		/*for j := range centroid[i] {
			print(centroid[i][j], "\n")
		}*/
	}
	clusteredData, err := kmeans(data, centroid, distanceFunction, threshold)

	return clusteredData, err
}
