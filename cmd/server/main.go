package main

import (
	"kaoluma/internal/server"
	"log"
	"net"
)

func main() {
	port := ":8080"
	listener, err := server.StartListener(port)
	if err != nil {
		log.Fatalln("failed to start listener:", err)
	}
	defer listener.Close()
	log.Println("server listening on port", port)

	server.AcceptLoop(listener, func(conn net.Conn) {
		log.Println("connection:", conn.RemoteAddr())
		server.HandleConnection(conn)
	})
}
