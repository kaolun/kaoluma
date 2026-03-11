package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {

		conn.SetReadDeadline(time.Now().Add(time.Second))

		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			} else {
				log.Println("Connection closed:", err)
				break
			}
		}

		fmt.Println("Recieved:", string(buf[:n]))

		fmt.Println(conn)
		fmt.Fprintf(conn, "Echo!\n")
		if err != nil {
			log.Println("error writing to client:", err)
			break
		}

	}
}
