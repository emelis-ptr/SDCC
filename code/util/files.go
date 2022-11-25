package util

import (
	"SDCC-project/code/mapreduce"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"testing"
)

// OpenEnv : apertura del file .env
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

// WriteClusters : scrive le informazioni dei cluster su un file json
func WriteClusters(clusters []mapreduce.Clusters, numPoint int, numMapper int, numReducer int, algo string) {
	fileName := algo + "-" + strconv.Itoa(numPoint) + "-" + strconv.Itoa(numMapper) + "-" + strconv.Itoa(numReducer) + ".txt"

	os.Mkdir(DirVolume+"/clusters", os.ModePerm)
	f, err := os.Create("/" + DirVolume + "/clusters/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	file, _ := json.MarshalIndent(clusters, "", " ")

	_, err = f.Write(file)
	if err != nil {
		log.Fatal("Errore scrittura file clusters", err)
	}
	fmt.Println("done")
}

// WriteBenchmark : scrive le informazioni del benchmark su un file txt
func WriteBenchmark(res testing.BenchmarkResult, numPoint int, numMapper int, numReducer int, algo string) {
	fileName := algo + "-" + strconv.Itoa(numPoint) + "-" + strconv.Itoa(numMapper) + "-" + strconv.Itoa(numReducer) + ".txt"

	os.Mkdir(DirVolume+"/benchmark", os.ModePerm)
	f, err := os.Create("/" + DirVolume + "/benchmark/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	file, _ := json.MarshalIndent(res, "", " ")

	_, err = f.WriteString(" ************************* ")
	_, err = f.WriteString("Algoritmo: " + algo)
	_, err = f.WriteString(" - Punti: " + strconv.Itoa(numPoint))
	_, err = f.WriteString("Mapper: " + strconv.Itoa(numMapper) + " - Reducer: " + strconv.Itoa(numReducer))
	_, err = f.Write(file)
	if err != nil {
		log.Fatal("Errore scrittura file benchmark", err)
	}
	fmt.Println("done")
}
