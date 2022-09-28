package main

/**
Nella fase di Reduce, avviene il calcolo dei nuovi centroidi: ciascun reducer in parallelo riceve in input
tutti i punti assegnati ad un determinato cluster e calcola il valore del centroide di quel cluster.
*/

func reduce(clusteredPoint []ClusteredPoint, numCluster int, distanceFunction DistanceFunction, threshold int) ([]ClusteredPoint, error) {

	centroid := newCentroid(clusteredPoint, numCluster)
	for i := range centroid {
		print(centroid[i], " centroid \n")
		/*for j := range centroid[i] {
			print(centroid[i][j], "\n")
		}*/
	}
	clusteredData, err := kmeansReduce(clusteredPoint, centroid, distanceFunction, threshold)

	return clusteredData, err
}

func kmeansReduce(data []ClusteredPoint, centroid []Point, distanceFunction DistanceFunction, threshold int) ([]ClusteredPoint, error) {
	var changes int
	for ii, p := range data {
		if closestCluster, _ := near(p, centroid, distanceFunction); closestCluster != p.ClusterNumber {
			changes++
			data[ii].ClusterNumber = closestCluster
		}
	}

	return data, nil
}

func newCentroid(clusteredPoint []ClusteredPoint, numCluster int) []Point {

	s := make([][]Point, numCluster)
	centroid := make([]Point, 0)

	var lenPoint int
	for i := 0; i < numCluster; i++ {
		for ii := range clusteredPoint {
			if (clusteredPoint[ii].ClusterNumber) == i {
				s[i] = append(s[i], clusteredPoint[ii].Point)
			}
			lenPoint = len(clusteredPoint[ii].Point)
		}

	}

	for _, i := range s {
		p := make([][]float64, lenPoint)

		for _, j := range i {
			for k := range j {
				p[k] = append(p[k], j[k])
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
		centroid = append(centroid, mean)
	}

	return centroid
}
