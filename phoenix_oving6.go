package main

import (
	"log"
	"net"
	"os/exec"
	"strconv"
	"time"
)

func broadcastUdp(addr string, counter int) {
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
		udpBroadcast.Write([]byte(strconv.Itoa(counter)))
		counter = counter + 1
		time.Sleep(1 * time.Second)
	}
}

func listenUdp(port string) {
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}
	udpListen, err := net.ListenUDP("udp", udpAddr)
	udpListen.SetDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	defer udpListen.Close()

	var buffer []byte
	count := 0
	for {
		n, ipAddr, err := udpListen.ReadFromUDP(buffer)
		if err != nil {
			//Kjør prosessen som en ny tråd
			go theProcess(count)
			return
		}
		count = strconv.Atoi(string(buffer[:n]))

		udpListen.SetDeadline(time.Now().Add(2 * time.Second))
	}
}

func theProcess(count int) {
	go broadcastUdp("129.241.187.255", count)

	log.Println("New process started")

	cmd := exec.Command("go run phoenix_oving6.go", "5") //WTF IS THIS
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	go listenUdp(":20010")
	for {

	}
}
