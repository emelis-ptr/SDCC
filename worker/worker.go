package main

import (
	"fmt"
	"log"
	"main/mapreduce"
	"main/util"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Worker is up")

	var conf util.Conf
	conf.ReadConf()

	api := new(mapreduce.API)
	//server := rpc.NewServer()
	err := rpc.Register(api)

	util.OpenEnv()

	num, _ := strconv.Atoi(os.Getenv("NUM_WORKER"))

	port := conf.PeerPort + rand.Intn(num)
	conf.PeerPort = port
	conf.PeerIP = util.GetOutboundIP()

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(conf.PeerPort))
	if err != nil {
		log.Fatal("Errore nella registrazione listener ", err)
	}

	log.Printf("Serving rpc sulla porta %s", strconv.Itoa(conf.PeerPort))
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Errore in serving : ", err)
	}
}
