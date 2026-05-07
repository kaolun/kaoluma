package client

import (
	"kaoluma/internal/core/protocol"
	"net"
)

type Client struct {
	Conn net.Conn

	Visibile bool
	JoinCode string

	UI *UIState

	SendQueue chan protocol.Packet
}

func RenderLoop(c *Client) {

}

func InputLoop(c *Client) {

}

type UIState struct {
	Online  bool
	Visbile bool

	Messages        []string
	PendingRequests []uint32
}
