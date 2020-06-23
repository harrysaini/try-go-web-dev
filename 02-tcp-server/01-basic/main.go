package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "Hello \n")

		fmt.Fprintln(conn, "How are you?")

		fmt.Println("Write now")

		for {

			reader := bufio.NewReader(os.Stdin)

			line, err := reader.ReadString('\n')
			if err != nil {
				log.Println(err)
				continue
			}

			if strings.Contains(line, "QUIT") {
				fmt.Println("QUITING")
				break
			}

			io.WriteString(conn, line)
		}

		conn.Close()

	}
}
