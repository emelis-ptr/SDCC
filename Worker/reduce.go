package main

//Reduce
/**
Nella fase di Reduce, avviene il calcolo dei nuovi centroidi: ciascun reducer in parallelo riceve in input
tutti i punti assegnati ad un determinato cluster e calcola il valore del centroide di quel cluster.
*/
func (a *API) Reduce(numCluster int, clusteredPoint *[]ClusteredPoint) error {

	_ = NewCentroid(*clusteredPoint, numCluster)

	return nil
}

// NewCentroid Determina i nuovi centroidi in base all'insieme dei punti del cluster
func NewCentroid(clusteredPoint []ClusteredPoint, numCluster int) []Point {

	s := make([][]Point, numCluster)
	centroid := make([]Point, 0)

	var lenPoint int
	for i := 0; i < numCluster; i++ {
		//Per ogni punto dell'insieme verifica se il valore del centroide assegnato precedentemente
		// corrisponde al valore del centroide, in modo tale da creare il cluster con l'insieme dei punti
		// che appartengono ad esso
		for ii := range clusteredPoint {
			if (clusteredPoint[ii].ClusterNumber) == i {
				s[i] = append(s[i], clusteredPoint[ii].Point)
			}
			lenPoint = len(clusteredPoint[ii].Point)
		}

	}

	//Calcola la media dei punti all'interno di un cluster per
	// determinare i nuovi centroidi
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
