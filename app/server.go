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
	"github.com/codecrafters-io/redis-starter-go/utils"
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
		Store: make(map[string]types.TValue),
	}

	for {
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
		commands := utils.RespParser(data)
		switch commands[0] {
		case "ping":
			if len(commands) != 1 {
				continue
			}

			router.PingHandler(conn)

		case "echo":
			if len(commands) != 2 {
				continue
			}

			router.EchoHandler(commands[1], conn)

		case "set":
			if len(commands) != 3 && len(commands) != 5 {
				continue
			}

			if len(commands) == 3 {
				router.NSetHandler(commands[1], commands[2], conn, global)
			} else if len(commands) == 5 && commands[3] == "px" {
				router.ESetHandler(commands[1], commands[2], commands[4], conn, global)
			}

		case "get":
			if len(commands) != 2 {
				continue
			}

			router.GetHandler(commands[1], conn, global)
		}
	}
}
