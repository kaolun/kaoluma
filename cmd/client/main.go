package main

import (
	"kaoluma/internal/client"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"log"
)

func main() {
	//rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	port := ":8080"

	conn, err := client.Connect(port)
	if err != nil {
		log.Fatal("connection failed:", err)
	}
	err = transport.InitiateHandshake(conn)
	if err != nil {
		log.Fatal("Handshake failed:", err)
	}

	c := &client.Client{
		Conn: conn,

		SendQueue: make(chan protocol.Packet, 100),
	}

	go client.ReadLoop(c)
	go client.SendLoop(c)
	go client.PingLoop(c)
	log.Println("cli is not designed for ease of use")

	go client.RenderLoop(c)
	go client.InputLoop(c)

	select {}
}
