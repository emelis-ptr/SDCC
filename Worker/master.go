package main

// Points Point con il numero cluster di appartenza
type Points struct {
	Point
}

type Cluster struct {
	Centroid   Centroid
	PointsData []Points
}

type Centroid struct {
	Index    int
	Centroid Point
}

// DistanceMethod Calcola la distanza tra due punti
type DistanceMethod func(first, second []float64) (float64, error)

/*func main() {
	numPoint, _ := strconv.Atoi(os.Getenv("numPoint"))
	numCentroid, _ := strconv.Atoi(os.Getenv("numCentroid"))
	numVector, _ := strconv.Atoi(os.Getenv("numVector"))
	numInterazioni, _ := strconv.Atoi(os.Getenv("numInterazioni"))

	print(numCentroid, "\n")
	print(numInterazioni, "\n")

	var err error
	var num int
	var c Conf
	c.readConf()

	num, _ = strconv.Atoi(os.Getenv("N"))
	clients := make([]*rpc.Client, num)

	for i := 0; i < num; i++ {
		port := c.PeerPort + i
		c.PeerPort = port
		clients[i], err = rpc.DialHTTP("tcp", "localhost:"+strconv.Itoa(port))
		if err != nil {
			fmt.Println("Errore di connessione , retring ... ")
		}
	}

	Master(clients, numPoint, numVector)
}

//Funzione eseguita dai thread figli per passare ai vari mapper la porzione di file da analizzare
//Il thread resta in attesa della risposta del mapper e successivamente la comunica attraverso un
//channel al main thread.

func rpcMap(point []Point, cli *rpc.Client, clusteredPoint []Points) {
	var reply []Points
	err := cli.Call("API.Mapper", point, &clusteredPoint)
	if err != nil {
		print("Error \n")
	}
	clusteredPoint <- reply
}

// Master Nel main viene effettuata solo la connessione con i worker per le chiamate rpc
func Master(cli []*rpc.Client, numPoint int, numVector int) {
	data := GeneratePoint(numPoint, numVector)
	clusteredPoint := CreateClusteredPoint(data)

	nodes := len(cli)
	c := make(chan map[string]int)

	chunk := len(clusteredPoint)

	j := 0
	for i := 0; i < nodes; i++ {
		go rpcMap(clusteredPoint[j].Point, cli[i], c)
		if i == (nodes-1) && j < chunk {
			i = 0
		}
		if j == (chunk - 1) {
			break
		}
		j++
	}
	print(j)
}*/
