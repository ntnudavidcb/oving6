package main

import{
	"fmt"
	"net"
}

var port int = 12020

func main(){
	// client

	BROADCAST_IPv4 := net.IPv4(255, 255, 255, 255)
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
	        IP:   BROADCAST_IPv4,
	        Port: port,
	})

	// server

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
	        IP:   net.IPv4(0, 0, 0, 0),
	        Port: port,
	})
	for {
	        data := make([]byte, 4096)
	        read, remoteAddr, err := socket.ReadFromUDP(data)
	}
}