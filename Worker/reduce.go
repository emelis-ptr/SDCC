package main

//Reduce
/**
Nella fase di Reduce, avviene il calcolo dei nuovi centroidi: ciascun reducer in parallelo riceve in input
tutti i punti assegnati ad un determinato cluster e calcola il valore del centroide di quel cluster.
*/
func Reduce(clusteredPoint *[]Cluster) ([]Centroid, error) {
	cp := NewCentroid(*clusteredPoint)

	return cp, nil
}

// NewCentroid Determina i nuovi centroidi in base all'insieme dei punti del cluster
func NewCentroid(clusters []Cluster) []Centroid {
	centroid := make([]Centroid, len(clusters))
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
