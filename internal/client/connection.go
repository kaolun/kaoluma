package client

import "net"

func Connect(port string) (net.Conn, error) {
	host := "localhost" + port
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
