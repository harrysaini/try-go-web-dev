package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	io.WriteString(conn, "Hello from TCP\n")
	conn.Close()
}
