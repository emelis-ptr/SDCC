package util

import (
	"SDCC-project/code/mapreduce"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
	"os"
	"strconv"
)

const (
	nameDir = "png"
	extFile = "png"
)

// Plot : rappresenta graficamente i punti e in centroidi di ogni cluster
func Plot(clusters []mapreduce.Clusters, numMapper int, numReducer int, numPoint int) {
	p := plot.New()

	p.Title.Text = "KMeans"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	for i := range clusters {
		for j := range clusters[i].PointsData {
			dataPoint := XY(len(clusters[i].PointsData), clusters[i].PointsData[j].Point)

			ss, err := plotter.NewScatter(dataPoint)
			if err != nil {
				panic(err)
			}
			ss.GlyphStyle.Color = plotutil.Color(i)
			ss.GlyphStyle.Shape = plotutil.Shape(i)
			ss.GlyphStyle.Radius = vg.Points(5)

			p.Add(ss)
		}

		dataCluster := XY(len(clusters), clusters[i].Centroid.Centroid)
		ss, err := plotter.NewScatter(dataCluster)
		if err != nil {
			panic(err)
		}
		ss.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0}
		ss.GlyphStyle.Shape = plotutil.Shape(i)
		ss.GlyphStyle.Radius = vg.Points(5)

		p.Add(ss)
	}

	os.Mkdir(DirVolume+"/"+nameDir, os.ModePerm)
	filename := DirVolume + "/" + nameDir + "/kmeans-" + strconv.Itoa(numPoint) + "-" + strconv.Itoa(numMapper) + "-" + strconv.Itoa(numReducer) + "." + extFile
	if err := p.Save(20*vg.Inch, 10*vg.Inch, filename); err != nil {
		panic(err)
	}
}

// XY Assegna ad x e y i valori del punto
func XY(numPoint int, observation []float64) plotter.XYs {
	pts := make(plotter.XYs, numPoint)
	for i := 0; i < numPoint; i++ {
		pts[i].X = observation[0]
		pts[i].Y = observation[1]
	}
	return pts
}
