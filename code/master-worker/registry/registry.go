package main

import (
	"SDCC-project/code/util"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type Registry struct {
	Peer util.Registration
}

var members = 0
var registration util.Registration

// RegisterMember : aggiunge un nuovo membro e ritorna in res la membership quando tutti sono stati registrati
func (registry *Registry) RegisterMember(arg util.Peer, res *util.Registration) error {

	registration.Index = members
	members++ // incrementa il numero di membri registrati

	registration.Peer = append(registration.Peer, arg)
	registry.Peer = registration
	*res = registration

	return nil
}

// RetrieveMember : recuper la membership
func (registry *Registry) RetrieveMember(bool bool, res *util.Registration) error {
	registration = registry.Peer

	if registration.Peer != nil {
		bool = true
	}
	*res = registration
	return nil
}

/*
*
Esecuzione del registry
*/
func main() {
	var err error

	fmt.Println("Registration service is up")

	util.OpenEnv()

	//read config
	var conf util.Conf
	conf.ReadConf(util.ConfPath)

	//expose api on open port
	err = rpc.Register(new(Registry))
	if err != nil {
		log.Fatalln("Error in register server name", err)
	}

	//init variables and signal handler for shutdown
	sigs := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//cancel context on shutdown
	go func() {
		<-sigs
		cancel()
	}()

	fmt.Println("Registration is starting...")

	//set up listening for incoming connection
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(conf.RegPort))
	if err != nil {
		log.Fatalln("Listening failed with error: ", err)
	}

	log.Printf("Serving rpc sulla porta %s", strconv.Itoa(conf.RegPort))
	log.Printf("Address reg %s", conf.RegIP)

	rpc.HandleHTTP()
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return http.Serve(lis, nil)
	})

	//close listener on shutdown
	g.Go(func() error {
		<-gCtx.Done()
		return lis.Close()
	})
	if err := g.Wait(); err != nil {
		fmt.Println("\nRegistration service shutdown")
	}
}
