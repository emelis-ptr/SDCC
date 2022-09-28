package main

// Observations is a slice of observations

//Point - Slice of Float64 Values
type Point []float64

// Add Summation of two vectors
func (observation Point) Add(otherObservation Point) {
	for ii, jj := range otherObservation {
		observation[ii] += jj
	}
}

// Mul Multiplication of a vector with a scalar
func (observation Point) Mul(scalar float64) {
	for ii := range observation {
		observation[ii] *= scalar
	}
}

// InnerProduct Dot Product of Two vectors
func (observation Point) InnerProduct(otherObservation Point) {
	for ii := range observation {
		observation[ii] *= otherObservation[ii]
	}
}

// OuterProduct Outer Product of two arrays
// TODO: Need to be tested
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
