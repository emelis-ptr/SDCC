package mapreduce

type Points struct {
	ClusterNumber Centroids
	Centroids     []Centroids
	Point         []float64
	Distance      float64
}

type Clusters struct {
	Centroid   Centroids
	PointsData []Points
	Changes    int
}

type Centroids struct {
	Index    int
	Centroid []float64
}
