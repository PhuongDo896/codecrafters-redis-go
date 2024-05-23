package main

import (
	"fmt"
	"log"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

const RESPONSE = "+PONG\r\n"

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}

	// for {
	// 	b := make([]byte, 1024)
	// 	_, err = conn.Read(b)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if strings.Contains(string(b), "PING") {
	// 		conn.Write([]byte(RESPONSE))
	// 	}
	// }
}

func handleConnection(conn net.Conn) {
	for {
		b := make([]byte, 1024)
		_, err := conn.Read(b)
		if err != nil {
			log.Println("DEBUGGING: ", err.Error())
		}

		if strings.Contains(string(b), "PING") {
			conn.Write([]byte(RESPONSE))
		}
	}
}
