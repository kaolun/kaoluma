package main

import (
	"kaoluma/internal/client"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"log"
	"time"
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
		UI: &client.UIState{
			Visible: false,
			Online:  false,
		},
		SendQueue: make(chan protocol.Packet, 100),
	}

	go client.ReadLoop(c)
	go client.SendLoop(c)
	go client.PingLoop(c)
	log.Println("cli is not designed for ease of use")

	time.Sleep(time.Second)
	c.Render()
	go c.InputLoop()

	select {}
}
