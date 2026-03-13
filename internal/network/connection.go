package network

import (
	"kaoluma/internal/core/protocol"
	"log"
	"net"
	"time"
)

// maybe something like
// handshake
// loop {
// read packets
// route what to do based on protocol case
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	//handshake? here later?

	for {
		err := conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		if err != nil {
			log.Println(err)
			continue
		}

		packet, err := protocol.DecodePacket(conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			log.Println("connection closed:", err)
			break
		}
		log.Println(packet)
		//dispatcher function

	}
}
