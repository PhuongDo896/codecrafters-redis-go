package router

import (
	"log"
	"net"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/types"
	"github.com/codecrafters-io/redis-starter-go/utils"
)

// 0 input, response pong
func PingHandler(s string, writer net.Conn) {
	if strings.Contains(s, "PING") {
		_, err := writer.Write(utils.Response("PONG"))
		if err != nil {
			log.Println("Error writing to connection: ", err.Error())
		}
	}
}

// 1 input, response same input
func EchoHandler(s string, writer net.Conn) {
	commands := utils.RespParser(s)
	if len(commands) != 2 {
		return
	}

	if commands[0] == "echo" {
		writer.Write(utils.Response(commands[1]))
	}
}

// 2 input, response OK, save key-value pair into global map
func SetHandler(s string, writer net.Conn, globalMap *types.GlobalMap) {
	commands := utils.RespParser(s)
	if len(commands) != 3 {
		return
	}

	if commands[0] == "set" {
		globalMap.Set(commands[1], commands[2])
		writer.Write(utils.Response("OK"))
	}
}

// 1 input, response value of key from global map
func GetHandler(s string, writer net.Conn, globalMap *types.GlobalMap) {
	commands := utils.RespParser(s)
	if len(commands) != 2 {
		return
	}

	if commands[0] == "get" {
		value := globalMap.Get(commands[1])
		writer.Write(utils.Response(value))
	}
}
