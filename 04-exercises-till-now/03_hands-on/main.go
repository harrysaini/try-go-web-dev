package main

import (
	"bufio"
	"fmt"
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

	scnr := bufio.NewScanner(conn)

	for scnr.Scan() {
		line := scnr.Text()
		fmt.Println(line)
	}

	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.")

	io.WriteString(conn, "Hello from TCP\n")
	conn.Close()
}
