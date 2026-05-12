package server

import (
	"crypto/rand"
	"kaoluma/internal/core/protocol"
	"math/big"
)

func (c *ClientInfo) GenerateJoincode(s *Server) error {
	code := generateCode()

	s.Mu.Lock()
	if c.JoinCode != "" {
		delete(s.CodeToClient, c.JoinCode)
	}
	c.JoinCode = code
	s.CodeToClient[code] = c
	s.Mu.Unlock()

	c.SendQueue <- protocol.Packet{
		Type: protocol.JoinCodeResponse,
		Data: []byte(code),
	}
	return nil
}

func generateCode() string {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, 6)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		b[i] = chars[n.Int64()]
	}
	return string(b)
}
