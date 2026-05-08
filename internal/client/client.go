package client

import (
	"kaoluma/internal/core/protocol"
	"net"
)

type Client struct {
	Conn net.Conn

	Visible  bool
	JoinCode string

	UI *UIState

	SendQueue chan protocol.Packet
}

type UIState struct {
	Online  bool
	Visible bool
	Logs    []string
}
