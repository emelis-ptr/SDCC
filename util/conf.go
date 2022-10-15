package util

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
)

type Conf struct {
	RegPort    int    `json:"reg_port"`
	MasterPort int    `json:"master_port"`
	PeerPort   int    `json:"peer_port"`
	RegIP      string `json:"reg_ip"`
	MasterIP   string `json:"master_ip"`
	PeerIP     string `json:"peer_ip"`
}

// ReadConf read config from json file
func (c *Conf) ReadConf() {
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		log.Fatalln("Configuration file cannot be open: ", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {

		log.Fatalln("Some error occurred while reading from config file: ", err)

	}
	err = json.Unmarshal(byteValue, c)
	if err != nil {

		log.Fatalln("Configuration file cannot be decoded: ", err)

	}
}

// GetOutboundIP /*Obtain peer ip
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()

}
