package client

import (
	"kaoluma/internal/core/protocol"
	"net"
	"sync"
	"time"
)

const Port = ":8080"

type Client struct {
	Conn     net.Conn
	LastPong time.Time
	Done     chan struct{}

	JoinCode string
	Mu       sync.RWMutex

	UI *UIState

	SendQueue chan protocol.Packet
}

type UIState struct {
	Online bool
	Logs   []string

	Mu sync.RWMutex
}
