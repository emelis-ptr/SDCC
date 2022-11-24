package master

import (
	"SDCC-project/code/algorithm"
	"SDCC-project/code/util"
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"time"
)

// Master : master che prende dal registry i worker che sono stati registrati e assegna loro
// un job da eseguire
func Master(numWorker int, numPoint int, numCentroid int, numMapper int, numReducer int, algo string, regIP string, regPort int, testing bool) {
	var registrations util.Registration

	var bools bool
	client, err := rpc.DialHTTP("tcp", regIP+":"+strconv.Itoa(regPort))

	time.Sleep(time.Second)
	// recupera i worker dal registry
	err = client.Call("Registry.RetrieveMember", bools, &registrations)
	for err != nil {
		log.Fatal("Error call registry", err.Error())
	}

	if !testing {
		for i := range registrations.Peer {
			log.Printf("Port %s", strconv.Itoa(registrations.Peer[i].Port))
			log.Printf("Address %s ", registrations.Peer[i].Address)
		}
	}
	clients := make([]*rpc.Client, len(registrations.Peer))
	calls := make([]*rpc.Call, len(registrations.Peer))

	// crea una connessione con i diversi worker
	for i := 0; i < len(registrations.Peer); i++ {
		clients[i], err = rpc.DialHTTP("tcp", registrations.Peer[i].Address+":"+strconv.Itoa(registrations.Peer[i].Port))
		if err != nil {
			fmt.Println("Errore di connessione , retring ... ")
		}
	}

	numClient := len(clients)
	// verifica che non ci siano worker che non sono stati registrati
	if numClient < numWorker {
		workerCrash := numWorker - numClient
		log.Fatalf("Errore! %d worker hanno subito un crash", workerCrash)
	}

	data := util.GeneratePoint(numPoint)
	points := util.CreateClusteredPoint(data)

	switch algo {
	case util.Llyod:
		fmt.Println(" ** Llyod **")
		//creazione di punti e centroidi
		centroids := algorithm.CreateInitValue(algo, numCentroid, points)
		algorithm.Llyod(numClient, numMapper, numReducer, points, centroids, clients, calls, testing)

	case util.Standard:
		fmt.Println(" ** Standard KMeans **")
		centroids := algorithm.CreateInitValue(algo, numCentroid, points)
		algorithm.StandardKMeans(numClient, numMapper, numReducer, points, centroids, clients, calls, testing)

	case util.KmeansPlusPlus:
		fmt.Println(" ** KMeans++ **")
		centroids := algorithm.InitCentroidKMeansPlusPlus(points)
		algorithm.KMeansPlusPlus(numClient, numCentroid, numMapper, numReducer, points, centroids, clients, calls, testing)
	}

	/*rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(portMaster))
	if err != nil {
		log.Fatal("Errore nella registrazione listener ", err)
		return
	}

	log.Printf("Serving rpc sulla porta %s", strconv.Itoa(portMaster))
	log.Printf("Address %d", portMaster)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Errore in serving : ", err)
		return
	}*/
}
