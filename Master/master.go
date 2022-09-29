package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/rpc"
	"os"
)

type API int

func main() {
	numPoint := os.Getenv("numPoint")
	numCentroid := os.Getenv("numCentroid")
	numVector := os.Getenv("numVector")
	numInterazioni := os.Getenv("numInterazioni")

	print(numCentroid + " " + numPoint + " " + numVector + " " + numInterazioni)

}

//Funzione eseguita dai thread figli per passare ai vari mapper la porzione di file da analizzare
//Il thread resta in attesa della risposta del mapper e successivamente la comunica attraverso un
//channel al main thread.

func rpc_map(part string, cli *rpc.Client, c chan map[string]int) {
	var reply map[string]int
	err := cli.Call("API.Mapper", part, &reply)
	if err != nil {
		return
	}
	c <- reply
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Master Nel main viene effettuata solo la connessione con i worker per le chiamate rpc
func Master(cli []*rpc.Client, numPoint int) {
	nodes := len(cli)
	c := make(chan map[string]int)

	j := 0
	for i := 0; i < nodes; i++ {
		//go rpc_map(text[j], cli[i], c)
	}
}
