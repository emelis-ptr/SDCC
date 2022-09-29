package main

/*func main() {

	numPoint := 100
	numCentroid := 10
	numVector := 2
	numInterazioni := 10

	if (numCentroid == 1) || (numPoint <= 0) || (numCentroid > numPoint) {
		return
	}

	data := GeneratePoint(numPoint, numVector)
	clusteredPoint := CreateClusteredPoint(data)

	centroid := InitCentroid(clusteredPoint, numCentroid, EuclideanDistance)

	var nameFile = "init"
	Scatter(clusteredPoint, centroid, len(clusteredPoint), nameFile)

	for i := 0; i < numInterazioni; i++ {
		// Best Distance for Iris is Canberra Distance

		clusteredPoints, err := Mapper(clusteredPoint, centroid, EuclideanDistance)

		centroid, err = Reduce(clusteredPoints, numCentroid)

		var nameFile = "init" + strconv.Itoa(i)
		Scatter(clusteredPoint, centroid, len(clusteredPoint), nameFile)

		if err != nil {
			log.Fatal(err)
		}
	}
}*/
