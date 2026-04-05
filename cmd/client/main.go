package main

import (
	"kaoluma/internal/client"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"log"
	"time"
)

func main() {
	port := ":8080"
	conn, err := client.Connect(port)
	if err != nil {
		log.Fatal("connection failed:", err)
	}
	err = transport.InitiateHandshake(conn)
	if err != nil {
		log.Fatal("Handshake failed:", err)
	}

	c := &protocol.Client{
		Conn: conn,

		SendQueue: make(chan protocol.Packet, 100),
	}

	go client.ReadLoop(c)
	go client.SendLoop(c)
	go client.PingLoop(c)

	time.Sleep(time.Minute)
}
