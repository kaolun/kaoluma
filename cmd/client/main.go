package main

import (
	"kaoluma/internal/client"
	"kaoluma/internal/core/protocol"
	"log"
	"time"
)

func main() {
	//rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	c := &client.Client{
		UI: &client.UIState{
			Online: false,
		},
		SendQueue: make(chan protocol.Packet, 100),
	}

	log.Print("\033[H\033[2J")
	log.Println("cli is designed for development")

	time.Sleep(time.Second * 3)
	c.Render()
	go c.InputLoop()

	select {}
}
