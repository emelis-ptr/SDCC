package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"log"
	"math/rand"
)

//Scatter Crea dei plot su uno spazio con i punti e centroidi
func Scatter(clusters []Cluster, nameFile string) {
	p := plot.New()
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	var err error

	for i := range clusters {
		red := rand.Intn(255)
		green := rand.Intn(255)
		blue := rand.Intn(255)

		for j := range clusters[i].PointsData {
			PlotPoints(p, clusters[i].PointsData[j].Point, len(clusters[i].PointsData), uint8(red), uint8(green), uint8(blue), 255)
		}
		PlotPoints(p, clusters[i].Centroid.Centroid, len(clusters), 22, 160, 133, 1)
	}

	err = p.Save(800, 500, "../Plot/"+nameFile+".png")
	if err != nil {
		log.Panic(err)
	}
}

// ScatterInit Crea dei plot su uno spazio con i punti e centroidi
func ScatterInit(clusteredPoint []Points, centroids []Centroid, nameFile string) {
	p := plot.New()
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	var err error

	for i := range clusteredPoint {
		PlotPoints(p, clusteredPoint[i].Point, len(clusteredPoint), 255, 128, 255, 255)
	}

	for i := 0; i < len(centroids); i++ {
		PlotPoints(p, centroids[i].Centroid, len(centroids), 22, 160, 133, 1)
	}

	err = p.Save(800, 500, "../Plot/"+nameFile+".png")
	if err != nil {
		log.Panic(err)
	}
}

func PlotPoints(p *plot.Plot, point Point, len int, r uint8, g uint8, b uint8, a uint8) {
	dataPoint := point.XY(len)
	lineData := point.XY(len)
	linePointsData := point.XY(len)

	s, err := plotter.NewScatter(dataPoint)
	if err != nil {
		log.Panic(err)
	}

	s.GlyphStyle.Color = color.RGBA{R: r, G: g, B: b}
	s.GlyphStyle.Radius = vg.Points(5)

	l, err := plotter.NewLine(lineData)
	if err != nil {
		log.Panic(err)
	}

	l.LineStyle.Width = vg.Points(5)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: b, A: a}

	lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
	if err != nil {
		log.Panic(err)
	}
	lpLine.Color = color.RGBA{G: g, A: a}
	lpPoints.Shape = draw.CircleGlyph{}
	lpPoints.Color = color.RGBA{R: r, A: a}
	lpPoints.GlyphStyle.Radius = vg.Points(5)

	p.Add(s, l, lpLine, lpPoints)
}
