package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// getFromFile Ottenere un insieme di punti attraverso un file
func getFromFile(nameFile string) [][]float64 {
	filePath, err := filepath.Abs(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	data := make([][]float64, len(lines))
	for ii, line := range lines {
		vector := strings.Split(line, ",")
		vector = vector[:len(vector)-1]
		floatVector := make([]float64, len(vector))
		for jj := range vector {
			floatVector[jj], err = strconv.ParseFloat(vector[jj], 64)
		}
		data[ii] = floatVector
	}

	return data
}
