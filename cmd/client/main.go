package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	//var i string
	//fmt.Scan(&i)
	data := []byte{
		1,
		0, 0, 0, 2,
		143, 67,
	}
	conn.Write(data)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(buf[:n]))

	conn.Close()
}
