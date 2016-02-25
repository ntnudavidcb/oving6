package main

import (
	"log"
	"net"
	"os/exec"
	"strconv"
	"time"
	"strings"
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
	udpListen, err := net.ListenUDP("udp4", udpAddr)
	udpListen.SetDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	defer udpListen.Close()
	buffer := make([]byte, 1024)
	count := 0
	for {
		log.Println("For udpListen")
		_, _, err := udpListen.ReadFromUDP(buffer)
		if err != nil {
			//Kjør prosessen som en ny tråd
			log.Println("Starting new process")
			go theProcess(count)
			return
		}
		//var i string = string(buffer[:n])

		readInt:=string(buffer)
		count,_ = strconv.Atoi(strings.Split(readInt, "\x00")[0])
		udpListen.SetDeadline(time.Now().Add(3 * time.Second))
	}
}

func theProcess(count int) {
	go broadcastUdp("192.168.56.255"+":30032", count)

	log.Println("New process started")

	cmd := exec.Command("gnome-terminal", "-x", "go", "run", "phoenix_oving6.go") //WTF IS THIS
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	go listenUdp(":30032")
	for {

	}
}
