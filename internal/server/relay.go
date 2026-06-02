package server

import (
	"fmt"
	"kaoluma/internal/core/protocol"
)

func (s *Server) allowed(from, to uint32) bool {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	fromClient := s.Clients[from]
	toClient := s.Clients[to]

	if toClient == nil {
		return false
	}

	return fromClient.PeerID == to && toClient.PeerID == from
}

func (s *Server) forward(sender *ClientInfo, p protocol.Packet) error {
	s.Mu.RLock()
	target := s.Clients[p.TargetID]
	s.Mu.RUnlock()

	if target == nil {
		fmt.Println("doesnt exist")
		return fmt.Errorf("targetID does not exist")
	}
	p.SenderID = sender.ID

	target.SendQueue <- p

	return nil
}
