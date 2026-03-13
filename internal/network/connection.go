package network

import (
	"net"
)

// maybe something like
// handshake
// loop {
// read packets
// route what to do based on protocol case
func HandleConnection(conn net.Conn) {
	defer conn.Close()

}
