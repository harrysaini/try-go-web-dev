package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

	method, url := parseRequest(conn)

	sendResponse(conn, method, url)

	defer conn.Close()

}

func parseRequest(conn net.Conn) (method string, url string) {

	scnr := bufio.NewScanner(conn)

	for scnr.Scan() {
		line := scnr.Text()

		// request line
		fileds := strings.Fields(line)

		method = fileds[0]
		url = fileds[1]

		break

	}

	return method, url

}

func sendResponse(conn net.Conn, method string, url string) {
	body := fmt.Sprintf("<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><title></title></head><body><strong>You called %s - %s</strong></body></html>", method, url)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
