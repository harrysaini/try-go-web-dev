package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go startServer()
	go startClient()

	// Addded wait group so wait for go routine to actually start server
	wg.Wait()
}

func startClient() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	count := 0

	data := fmt.Sprintf("Hello world - %d", count)

	fmt.Println("CLIENT", "SENT", data)

	fmt.Fprintln(conn, data)

	scnr := bufio.NewScanner(conn)

	for scnr.Scan() {

		line := scnr.Text()

		fmt.Println("CLIENT", "RECEIVED", line)

		count++
		fmt.Fprintln(conn, fmt.Sprintf("Hello world - %d", count))
	}

	wg.Done()
}

func startServer() {
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

		go func(conn net.Conn) {
			reader := bufio.NewReader(conn)

			for {
				line, err := reader.ReadString('\n')

				if err != nil {
					log.Println(err)
					continue
				}

				fmt.Println("SERVER", "RECIEVED", line)

				time.Sleep(2 * time.Second)

				fmt.Fprint(conn, "echo from server: ", line)
			}

		}(conn)

	}

}
