package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

const (
	RESPONSE = "+PONG\r\n"
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

	for {
		// TODO: this block until new connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
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

		if strings.Contains(string(input), "PING") {
			conn.Write([]byte(RESPONSE))
		}

		commands := respParser(string(input))
		if len(commands) != 2 {
			continue
		}

		if commands[0] == "echo" {
			conn.Write(response(commands[1]))
		}
	}
}

func respParser(resp string) []string {
	// normalize input
	resp = strings.ToLower(resp)

	re := regexp.MustCompile(`\$([0-9]+)\r\n(.+?)\r\n`)
	match := re.FindAllString(resp, -1)

	command := make([]string, 0)

	for _, m := range match {
		word, ok := processBulkString(m)
		if ok {
			command = append(command, word)
		}
	}

	return command
}

// 1st bit: $
// 2nd bit: word length
// 3rd and 4th bit: \r\n
// 2 last bits: \r\n
// => total bits = word length + 6
func processBulkString(s string) (string, bool) {
	wordLen, err := strconv.Atoi(string(s[1]))
	if err != nil {
		return "", false
	}

	if len(s) != wordLen+6 {
		return "", false
	}

	return s[4 : 4+wordLen], true
}

func response(s string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", s))
}
