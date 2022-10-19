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
	"strconv"
	"time"
)

func main() {
	fmt.Println("Worker is up")

	var reply util.Registration
	var conf util.Conf
	conf.ReadConf()

	api := new(mapreduce.API)
	err := rpc.Register(api)

	util.OpenEnv()

	rand.Seed(987654321 * time.Now().UnixNano())
	n := rand.Intn(100)
	port := conf.PeerPort + n
	conf.PeerPort = port
	conf.PeerIP = util.GetOutboundIP()

	client, err := rpc.DialHTTP("tcp", conf.RegIP+":"+strconv.Itoa(conf.RegPort))

	peer := &util.Peer{Port: conf.PeerPort, Address: conf.PeerIP}

	time.Sleep(time.Second)
	err = client.Call("Registry.RegisterMember", peer, &reply)
	for err != nil {
		log.Fatalln("Error call registry", err.Error())
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(conf.PeerPort))
	if err != nil {
		log.Fatal("Errore nella registrazione listener ", err)
		return
	}

	log.Printf("Serving rpc sulla porta %s", strconv.Itoa(conf.PeerPort))
	log.Printf("Address %s", conf.PeerIP)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Errore in serving : ", err)
		return
	}

}
