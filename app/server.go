package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sync"

	// Uncomment this block to pass the first stage
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/router"
	"github.com/codecrafters-io/redis-starter-go/types"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	globalMap := &types.GlobalMap{
		Mu:    sync.Mutex{},
		Store: make(map[string]string),
	}

	for {
		// TODO: this block until new connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn, globalMap)
	}

}

func handleConnection(conn net.Conn, global *types.GlobalMap) {
	defer conn.Close()

	for {
		input := make([]byte, 1024)
		// TODO: this block until new line is sent by client
		_, err := conn.Read(input)
		if errors.Is(err, io.EOF) {
			log.Println("Connection closed by client")
			break
		} else if err != nil {
			log.Println("DEBUGGING: ", err.Error())
			break
		}

		//	router
		data := string(input)
		router.PingHandler(data, conn)
		router.EchoHandler(data, conn)
		router.SetHandler(data, conn, global)
		router.GetHandler(data, conn, global)
	}
}
