package main

import (
	"log"
	"net"
	"time"
)

func broadcastUdp(addr string) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	udpBroadcast, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer udpBroadcast.Close()

	for {
		udpBroadcast.Write([]byte("Not master"))
		time.Sleep(1000 * time.Millisecond)
		log.Println("Hei")
	}
}

func ipInList(ipAddr string, ipList []string) bool {
	for _, b := range ipList {
		if b == ipAddr {
			return true
		}
	}
	return false
}

func main() {
	doneChannel := make(chan bool, 1)
	ipListChannel := make(chan []string, 1)
	doneChannel <- true
	broadcastAddr := "129.241.187.255:20060"

	go broadcastUdp(broadcastAddr)

	log.Println(<-ipListChannel)
	<-doneChannel
}
