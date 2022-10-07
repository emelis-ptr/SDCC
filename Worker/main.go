package main

/*func maina() {

	numPoint, _ := strconv.Atoi(os.Getenv("NUMPOINT"))
	numCentroid, _ := strconv.Atoi(os.Getenv("NUMCENTROID"))
	numVector, _ := strconv.Atoi(os.Getenv("NUMVECTOR"))
	maxInteration, _ := strconv.Atoi(os.Getenv("MAXINTERAZIONI"))

	if (numCentroid == 1) || (numPoint <= 0) || (numCentroid > numPoint) {
		return
	}

	data := GeneratePoint(numPoint, numVector)
	points := CreateClusteredPoint(data)
	centroids := InitCentroid(points, numCentroid, EuclideanDistance)

	var firstNameFile = "kmeans-" + strconv.Itoa(numPoint) + "-" + strconv.Itoa(numCentroid)
	ScatterInit(points, centroids, firstNameFile)

	for i := 0; i < maxInteration; i++ {
		cluster := make([]Clusters, numCentroid)

		for i := range centroids {
			cluster[i].Centroids.Index = i
			cluster[i].Centroids.Centroids = centroids[i].Centroids
		}

		for j := range points {
			cluster, points[j], _ = Mapper(&points[j], &cluster)
		}

		centroids, _ = Reduce(cluster, &centroids)

		var nameFile = firstNameFile + "-" + strconv.Itoa(i)
		Scatter(cluster, nameFile)

	}

}*/
