package main

import (
	"bufio"
	"bytes"
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
		if line == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}

	fmt.Println("Code got here.")

	writer := bytes.NewBuffer(*new([]byte))

	io.WriteString(writer, "I see you connected.\n")
	io.WriteString(writer, "Hello from TCP\n")

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")

	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(writer.String()))

	fmt.Fprint(conn, "Content-Type: text/plain\r\n")

	fmt.Fprint(conn, "\r\n")

	fmt.Fprint(conn, writer.String())

	conn.Close()
}
