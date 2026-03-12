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
	data := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 44, 32, 119, 104, 97, 116, 115, 32, 112, 111, 112, 112, 105, 110}
	conn.Write(data)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(buf[:n]))

	conn.Close()
}
