package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
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
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	scnr := bufio.NewScanner(conn)

	for scnr.Scan() {
		line := scnr.Text()
		fmt.Println(line)

		fmt.Fprintln(conn, "You sent: ", line)
	}
}
