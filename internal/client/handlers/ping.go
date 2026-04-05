package handlers

import (
	"kaoluma/internal/core/protocol"
	"log"
)

func HandlePong(packet protocol.Packet) error {
	log.Println("ping success")
	return nil
}
