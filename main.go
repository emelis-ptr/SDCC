package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	irisData := getFromFile("data/iris.csv")
	print(irisData, "\n")
	data := generatePoint(100, 2)

	threshold := 10
	numCentroid := 60
	// Best Distance for Iris is Canberra Distance
	clusteredPoints, err := mapper(data, numCentroid, 2, EuclideanDistance, threshold)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(clusteredPoints); i++ {
		print(clusteredPoints[i].Point, " ", clusteredPoints[i].ClusterNumber, "\n")
	}
	print(clusteredPoints, "\n")

	red, _ := reduce(clusteredPoints, numCentroid, EuclideanDistance, threshold)
	for i := 0; i < len(red); i++ {
		print(red[i].Point, " ", red[i].ClusterNumber, "\n")
		/*for j := range red[i].Point {
			print(red[i].Point[j], "\n")
		}*/
	}
	print(red, " red \n")

}

func generatePoint(numPoint int, numVector int) [][]float64 {
	data := make([][]float64, numPoint)
	for i := 0; i < numPoint; i++ {
		for j := 0; j < numVector; j++ {
			data[i] = append(data[i], rand.Float64())
		}
	}
	return data
}

func getFromFile(nameFile string) [][]float64 {
	filePath, err := filepath.Abs(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	data := make([][]float64, len(lines))
	for ii, line := range lines {
		vector := strings.Split(line, ",")
		vector = vector[:len(vector)-1]
		floatVector := make([]float64, len(vector))
		for jj := range vector {
			floatVector[jj], err = strconv.ParseFloat(vector[jj], 64)
		}
		data[ii] = floatVector
	}

	return data
}
