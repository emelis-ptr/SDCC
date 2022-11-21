package util

import (
	"SDCC-project/code/mapreduce"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

const (
	nameDir = "png"
	extFile = "png"
)

func Plot(points []mapreduce.Points) {
	p := plot.New()

	p.Title.Text = "KMeans"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	//p. = color.RGBA{R: 255, B: 255, A: 255}
	for i := range points {
		dataPoint := XY(len(points), points[i].Point)

		ss, err := plotter.NewScatter(dataPoint)
		if err != nil {
			panic(err)
		}
		ss.Color = color.Black
		p.Add(ss)
	}

	filename := "./doc/plot/kmeans.png"
	if err := p.Save(10*vg.Inch, 10*vg.Inch, filename); err != nil {
		panic(err)
	}
}

/*//Scatter Crea dei plot su uno spazio con i punti e centroidi
func Scatter(clusters []mapreduce.Clusters, nameFile string) {
	p := plot.New()
	p.Title.Text = "KMeans"
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

	pathFile := "./doc/plot/" + nameDir + "/" + nameFile + "." + extFile
	err = p.Save(1200, 800, pathFile)
	if err != nil {
		log.Panic(err)
	}
}

func PlotPoints(p *plot.Plot, point []float64, len int, r uint8, g uint8, b uint8, a uint8) {
	dataPoint := XY(len, point)
	lineData := XY(len, point)
	linePointsData := XY(len, point)

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
*/

// XY Assegna ad x e y i valori del punto
func XY(numPoint int, observation []float64) plotter.XYs {
	pts := make(plotter.XYs, numPoint)
	for i := 0; i < numPoint; i++ {
		pts[i].X = observation[0]
		pts[i].Y = observation[1]
	}
	return pts
}
