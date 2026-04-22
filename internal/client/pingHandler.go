package client

import (
	"kaoluma/internal/core/protocol"
	"log"
)

func handlePong(packet protocol.Packet) error {
	log.Println("ping success")
	return nil
}
