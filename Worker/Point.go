package main

import (
	"gonum.org/v1/plot/plotter"
)

//Point - Slice of Float64 Values
type Point []float64

//Len dimensione del punto
func (observation Point) Len() int {
	return len(observation)
}

//XY Assegna ad x e y i valori del punto
func (observation Point) XY(numPoint int) plotter.XYs {
	pts := make(plotter.XYs, numPoint)

	for i := 0; i < numPoint; i++ {
		pts[i].X = observation[0]
		pts[i].Y = observation[1]
	}
	return pts
}

//XY Assegna ad x e y i valori del punto
func XYFloat(x int, y float64, observation int) plotter.XYs {
	pts := make(plotter.XYs, observation)

	for i := 0; i < observation; i++ {
		pts[i].X = float64(x)
		pts[i].Y = y
	}
	return pts
}

// Add Somma di due vettori
func (observation Point) Add(otherObservation Point) {
	for ii, jj := range otherObservation {
		observation[ii] += jj
	}
}

// Mul Moltiplicazione di un vettore con uno scalare
func (observation Point) Mul(scalar float64) {
	for ii := range observation {
		observation[ii] *= scalar
	}
}

// InnerProduct Dot Prodotto tra due vettori
func (observation Point) InnerProduct(otherObservation Point) {
	for ii := range observation {
		observation[ii] *= otherObservation[ii]
	}
}

// OuterProduct Outer Prodotto tra due array
func (observation Point) OuterProduct(otherObservation Point) [][]float64 {
	result := make([][]float64, len(observation))
	for ii := range result {
		result[ii] = make([]float64, len(otherObservation))
	}
	for ii := range result {
		for jj := range result[ii] {
			result[ii][jj] = observation[ii] * otherObservation[jj]
		}
	}
	return result
}
