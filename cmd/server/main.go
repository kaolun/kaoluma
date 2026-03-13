package main

import (
	"kaoluma/internal/network"
	"log"
	"net"
)

func main() {
	port := ":8080"
	listener, err := network.StartListener(port)
	if err != nil {
		log.Fatalln("failed to start listener:", err)
	}
	defer listener.Close()
	log.Println("server listening on port:", port)

	network.AcceptLoop(listener, func(conn net.Conn) {
		log.Println("handling handling handling")
		network.HandleConnection(conn)
	})
}
