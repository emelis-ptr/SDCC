package util

import (
	"github.com/joho/godotenv"
	"log"
	"main/mapreduce"
	"os"
	"strconv"
)

const (
	nameFile = "kmeans"
)

//Open file .env
func OpenEnv() {
	//Apertura file.
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("failed to open")

	}
	err = godotenv.Load(file.Name())
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func writeFile(cluster []mapreduce.Clusters) {
	// create file
	f, err := os.Create("./logs/" + nameFile + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file
	defer f.Close()

	for _, line := range cluster {
		_, err = f.WriteString(strconv.Itoa(line.Centroid.Index) + " - ")
		for i := range line.Centroid.Centroid {
			_, err = f.WriteString(strconv.FormatFloat(line.Centroid.Centroid[i], 'f', 5, 64) + " ")
		}
		_, err = f.WriteString("\n")
		for i := range line.PointsData {
			for j := range line.PointsData[i].Point {
				_, err = f.WriteString(strconv.FormatFloat(line.PointsData[i].Point[j], 'f', 5, 64) + " ")
			}
			_, err = f.WriteString(" - ")
		}
		_, err = f.WriteString("\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
