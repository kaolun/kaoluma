package main

import (
	"kaoluma/internal/core/protocol"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	packet := protocol.Packet{
		Type: 1,
		Data: []byte{143, 67, 67, 67, 67, 143, 143},
	}
	data := protocol.EcodePacket(packet)
	conn.Write(data)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(buf[:n]))

	conn.Close()
}
