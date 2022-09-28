package main

import (
	"log"
	"strconv"
)

func main() {

	numPoint := 100
	numCentroid := 10
	numVector := 2
	numInterazioni := 10

	if (numCentroid == 1) || (numPoint <= 0) || (numCentroid > numPoint) {
		return
	}

	data := generatePoint(numPoint, numVector)
	clusteredPoint := createClusteredPoint(data)

	centroid := initCentroid(clusteredPoint, numCentroid, EuclideanDistance)

	var nameFile = "init"
	scatter(clusteredPoint, centroid, len(clusteredPoint), nameFile)

	for i := 0; i < numInterazioni; i++ {
		// Best Distance for Iris is Canberra Distance
		clusteredPoints, err := mapper(clusteredPoint, centroid, EuclideanDistance)

		centroid, err = reduce(clusteredPoints, numCentroid)

		var nameFile = "init" + strconv.Itoa(i)
		scatter(clusteredPoint, centroid, len(clusteredPoint), nameFile)

		if err != nil {
			log.Fatal(err)
		}
	}
}
