package main

import (
	"kaoluma/internal/client"
	"kaoluma/internal/core/transport"
	"log"
)

func main() {
	port := "8080"
	conn, err := client.Connect(port)
	if err != nil {
		log.Fatal("connection failed:", err)
	}

	err = transport.InitiateHandshake(conn)
	if err != nil {
		log.Fatal("CHECK VERSION: Handshake failed:", err)
	}
}
