package server

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"log"
	"net"
	"time"
)

func HandleConnection(s *Server, conn net.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	err := transport.PerformHandshake(conn)
	if err != nil {
		log.Println("handshake failed:", err)
		return
	}
	s.Mu.Lock()
	s.NextClientID++
	client := &ClientInfo{
		ID:        s.NextClientID,
		Conn:      conn,
		SendQueue: make(chan protocol.Packet, 100),
	}
	go SendLoop(client)
	s.Clients[client.ID] = client
	s.Mu.Unlock()

	for id, client := range s.Clients {
		log.Printf("id: %d, conn: %v\n", id, client.Conn)
	}

	for {
		err := conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		if err != nil {
			log.Println(err)
			continue
		}

		packet, err := protocol.DecodePacket(client.Conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			log.Println("connection closed:", err)
			s.Mu.Lock()
			delete(s.Clients, client.ID)
			s.Mu.Unlock()
			break
		}
		err = dispatch(s, client, packet)
		if err != nil {
			log.Println("error handling packet", err)
		}

	}
}
