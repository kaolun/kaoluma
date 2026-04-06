package server

import (
	"log"
	"net"
)

func StartListener(port string) (net.Listener, error) {
	host := "localhost" + port
	ln, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	return ln, nil
}

func AcceptLoop(s *Server, listener net.Listener, handleFunc func(s *Server, conn net.Conn)) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleFunc(s, conn)

	}
}
