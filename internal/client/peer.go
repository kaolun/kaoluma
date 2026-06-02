package client

import (
	"encoding/binary"
	"kaoluma/internal/core/protocol"
)

func (c *Client) RequestPeerConnection(code string) error {
	c.SendQueue <- protocol.Packet{
		TargetID: 0,
		Type:     protocol.RequestPeerConnection,
		Data:     []byte(code),
	}
	c.Log("Peer connection requested:" + code)
	return nil
}
func (c *Client) HandleConnectionPeerAccept(p *protocol.Packet) error {
	c.PeerID = binary.BigEndian.Uint32(p.Data)
	c.Log("Peer connection Accepted")
	return nil
}
func (c *Client) HandleRejectPeerConnection(p *protocol.Packet) error {
	c.Log("Peer connection Rejected")
	return nil
}
func (c *Client) HandleNotifyPeerConnection(p *protocol.Packet) error {
	c.Mu.Lock()
	c.PeerID = binary.BigEndian.Uint32(p.Data)
	c.Mu.Unlock()
	c.Log("Peer connected")
	return nil
}

func (c *Client) RequestPeerDisconnect() error {
	c.SendQueue <- protocol.Packet{
		TargetID: 0,
		Type:     protocol.DisconnectPeerConnection,
	}
	c.Log("Peer disconnect requested")
	return nil
}
func (c *Client) HandlePeerDisconnectConfirm(p *protocol.Packet) error {
	c.Mu.Lock()
	c.PeerID = 0
	c.Mu.Unlock()
	c.Log("Peer successfully disconnected")
	return nil
}
func (c *Client) HandleNotifyPeerDisconnect() error {
	c.Mu.Lock()
	c.PeerID = 0
	c.Mu.Unlock()
	c.Log("Peer disconnected")
	return nil
}
