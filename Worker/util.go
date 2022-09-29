package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type Algorithm int

type Req struct {
	P         int
	Timestamp []int
	IP        string
	Port      int
}

type Peer struct {
	IP   string
	Port int
}

type Registration_reply struct {
	Peer  []Peer
	Alg   Algorithm
	Index int
	Mask  []int
}

type Conf struct {
	RegPort    int    `json:"reg_port"`
	MasterPort int    `json:"master_port"`
	PeerPort   int    `json:"peer_port"`
	RegIP      string `json:"reg_ip"`
	MasterIP   string `json:"master_ip"`
	PeerIP     string `json:"peer_ip"`
}

func (c *Conf) readConf() {
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		log.Fatalln("Configuration file cannot be open: ", err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalln("File cannot be close")
		}
	}(jsonFile)

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln("Some error occurred while reading from config file: ", err)
	}
	err = json.Unmarshal(byteValue, c)
	if err != nil {
		log.Fatalln("Configuration file cannot be decoded: ", err)
	}
}

func findIndex(list []Peer, elem Peer) int {
	for i := range list {
		if list[i] == elem {
			return i
		}
	}
	return -1
}

func InitLogger(name string) (*log.Logger, error) {
	logFile, err := os.OpenFile(
		fmt.Sprintf("../logs/%v.log", name),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		return nil, err
	}
	my_log := log.New(logFile, "", log.LstdFlags)
	return my_log, nil
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()

}

func getPublicIP() (string, error) {
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
