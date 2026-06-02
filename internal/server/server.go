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

	CodeToClient map[string]*ClientInfo

	Mu sync.RWMutex
}

type ClientInfo struct {
	ID        uint32
	Conn      net.Conn
	SendQueue chan protocol.Packet
	Done      chan struct{}

	JoinCode string
	PeerID   uint32
	LastSeen time.Time
}

func NewServer() *Server {
	return &Server{
		Clients: make(map[uint32]*ClientInfo),

		CodeToClient: make(map[string]*ClientInfo),
	}
}
