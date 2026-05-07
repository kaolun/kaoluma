package server

import (
	"kaoluma/internal/core/protocol"
	"net"
	"sync"
	"time"
)

type Server struct {
	Clients      map[uint32]*ClientInfo
	NextClientID uint32

	PeerConnections map[uint32]map[uint32]bool
	CodeToClient    map[string]*ClientInfo

	Mu sync.RWMutex
}

type ClientInfo struct {
	ID        uint32
	Conn      net.Conn
	SendQueue chan protocol.Packet

	Visible  bool
	JoinCode string
	LastSeen time.Time
}

func NewServer() *Server {
	return &Server{
		Clients: make(map[uint32]*ClientInfo),

		PeerConnections: make(map[uint32]map[uint32]bool),
		CodeToClient:    make(map[string]*ClientInfo),
	}
}
