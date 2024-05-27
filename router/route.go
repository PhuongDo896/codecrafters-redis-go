package router

import (
	"log"
	"net"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/types"
	"github.com/codecrafters-io/redis-starter-go/utils"
)

// 0 input, response pong
func PingHandler(writer net.Conn) {
	writer.Write(utils.StatusResponse("PONG"))
}

// 1 input, response same input
func EchoHandler(data string, writer net.Conn) {
	writer.Write(utils.BulkStringResponse(data))
}

// 2 input, response OK, save key-value pair into global map
func NSetHandler(key, value string, writer net.Conn, globalMap *types.GlobalMap) {
	globalMap.NSet(key, value)
	writer.Write(utils.StatusResponse("OK"))
}

// 1 input, response value of key from global map
func GetHandler(key string, writer net.Conn, globalMap *types.GlobalMap) {
	value := globalMap.Get(key)
	if value == "" {
		writer.Write(utils.NullResponse())
	} else {
		writer.Write(utils.BulkStringResponse(value))
	}
}

// 4 input, response OK, save key-value pair into global map, delete it after expire time passed
func ESetHandler(key, value string, expireTime string, writer net.Conn, globalMap *types.GlobalMap) {
	convertTime, err := strconv.ParseInt(expireTime, 10, 64)
	if err != nil {
		log.Fatalf("Cannot convert expire time to int64: %v", err)
	}

	globalMap.ESet(key, value, convertTime)
	writer.Write(utils.StatusResponse("OK"))
}
