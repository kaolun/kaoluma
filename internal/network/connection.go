package network

import (
	"kaoluma/internal/core/protocol"
	"log"
	"net"
)

// maybe something like
// handshake
// loop {
// read packets
// route what to do based on protocol case
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	Packet, err := protocol.DecodePacket(conn)
	if err != nil {
		log.Println("failed to decode packet: ", err)
		return
	}
	log.Println(Packet)

}
