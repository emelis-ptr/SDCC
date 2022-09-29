package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"log"
)

//scatter Crea dei plot su uno spazio con i punti e centroidi
func scatter(clusteredPoint []ClusteredPoint, centroids []Point, numPoint int, nameFile string) {
	p := plot.New()
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	var err error

	for i := range clusteredPoint {
		dataPoint := clusteredPoint[i].Point.XY(numPoint)
		lineData := clusteredPoint[i].Point.XY(numPoint)
		linePointsData := clusteredPoint[i].Point.XY(numPoint)

		s, err := plotter.NewScatter(dataPoint)
		if err != nil {
			log.Panic(err)
		}

		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(5)

		l, err := plotter.NewLine(lineData)
		if err != nil {
			log.Panic(err)
		}
		l.LineStyle.Width = vg.Points(5)
		l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
		l.LineStyle.Color = color.RGBA{B: 255, A: 255}

		lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
		if err != nil {
			log.Panic(err)
		}
		lpLine.Color = color.RGBA{G: 255, A: 255}
		lpPoints.Shape = draw.CircleGlyph{}
		lpPoints.Color = color.RGBA{R: 255, A: 255}
		lpPoints.GlyphStyle.Radius = vg.Points(5)

		p.Add(s, l, lpLine, lpPoints)

	}

	for i := 0; i < len(centroids); i++ {
		centroidPoint := centroids[i].XY(len(centroids))
		lineData := centroids[i].XY(len(centroids))
		linePointsData := centroids[i].XY(len(centroids))

		s, err := plotter.NewScatter(centroidPoint)
		if err != nil {
			log.Panic(err)
		}

		s.Color = color.RGBA{G: 255, A: 255}
		s.GlyphStyle.Color = color.RGBA{G: 255, A: 255}
		s.GlyphStyle.Radius = vg.Points(5)

		l, err := plotter.NewLine(lineData)
		if err != nil {
			log.Panic(err)
		}
		l.LineStyle.Width = vg.Points(5)
		l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
		l.LineStyle.Color = color.RGBA{G: 255, A: 255}

		lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
		if err != nil {
			log.Panic(err)
		}
		lpLine.Color = color.RGBA{G: 255, A: 255}
		lpPoints.Shape = draw.CircleGlyph{}
		lpPoints.Color = color.RGBA{G: 255, A: 255}
		lpPoints.GlyphStyle.Radius = vg.Points(5)

		p.Add(s, l, lpLine, lpPoints)

		p.Add(s)
	}

	err = p.Save(800, 500, "plots/"+nameFile+".png")
	if err != nil {
		log.Panic(err)
	}
}
