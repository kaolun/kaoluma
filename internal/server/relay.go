package server

import (
	"fmt"
	"kaoluma/internal/core/protocol"
)

func (s *Server) allowed(from, to uint32) bool {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	if s.PeerConnections == nil {
		return false
	}
	if s.PeerConnections[from] == nil {
		return false
	}

	return s.PeerConnections[from][to]
}

func (s *Server) forward(sender *protocol.Client, p protocol.Packet) error {
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
