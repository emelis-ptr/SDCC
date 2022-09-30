package main

import (
	"log"
	"strconv"
)

func main() {

	numPoint := 100
	numCentroid := 5
	numVector := 2
	numInterazioni := 5
	print(numInterazioni)

	if (numCentroid == 1) || (numPoint <= 0) || (numCentroid > numPoint) {
		return
	}

	data := GeneratePoint(numPoint, numVector)
	clusteredPoint := CreateClusteredPoint(data)
	centroid := InitCentroid(clusteredPoint, numCentroid, EuclideanDistance)

	var firstNameFile = "kmeans-" + strconv.Itoa(numPoint) + "-" + strconv.Itoa(numCentroid)
	ScatterInit(clusteredPoint, centroid, firstNameFile)

	for i := 0; i < numInterazioni; i++ {

		clusters, err := Mapper(centroid, &clusteredPoint)
		centroid, err = Reduce(&clusters)

		var nameFile = firstNameFile + "-" + strconv.Itoa(i)
		Scatter(clusters, nameFile)

		if err != nil {
			log.Fatal(err)
		}

	}

}
