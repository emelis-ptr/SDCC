package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

type API int

type Data struct {
	clusteredPoint []ClusteredPoint
	point          []Point
}

func main() {
	var c Conf
	Worker(c)
}

func Worker(c Conf) {
	var api = new(API)
	n := rand.Intn(1000)

	fmt.Println("Peer client is up")

	c.readConf()

	port := c.PeerPort + n

	c.PeerPort = port
	c.PeerIP = GetOutboundIP()

	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Errore nella registrazione ", err)
	}

	//Consumer in attesa di ricevere chiamate rpc
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(c.PeerPort))
	if err != nil {
		log.Fatal("Errore nella registrazione listener ", err)
	}

	log.Printf("Serving rpc sulla porta %s", strconv.Itoa(c.PeerPort))
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Errore in serving : ", err)
	}

}
