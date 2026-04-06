package server

import (
	"kaoluma/internal/core/protocol"
)

type Server struct {
	Clients       map[uint32]*protocol.Client
	IncomingFiles map[uint32]*protocol.IncomingFile
	NextClientID  uint32
}

func NewServer() *Server {
	return &Server{
		Clients:       make(map[uint32]*protocol.Client),
		IncomingFiles: make(map[uint32]*protocol.IncomingFile),
	}
}
