package server

import (
	"kaoluma/internal/core/protocol"
	"sync"
)

type Server struct {
	Clients      map[uint32]*protocol.Client
	NextClientID uint32

	PeerConnections map[uint32]map[uint32]bool
	CodeToClient    map[string]*protocol.Client

	Mu sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		Clients: make(map[uint32]*protocol.Client),

		PeerConnections: make(map[uint32]map[uint32]bool),
		CodeToClient:    make(map[string]*protocol.Client),
	}
}
