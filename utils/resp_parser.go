package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RespParser(resp string) []string {
	// normalize input
	resp = strings.ToLower(resp)

	re := regexp.MustCompile(`\$([0-9]+)\r\n(.+?)\r\n`)
	match := re.FindAllStringSubmatch(resp, -1)

	commands := make([]string, 0)

	for _, m := range match {
		word, ok := processBulkString(m...)
		if ok {
			commands = append(commands, word)
		}
	}

	return commands
}

// for each submatch, there's 3 elements
// 1st elem: whole submatch
// 2nd elem: ([0-9]+) group
// 3rd elem: (.+?) group
func processBulkString(s ...string) (string, bool) {
	if len(s) != 3 {
		return "", false
	}

	wordLen, err := strconv.Atoi(string(s[1]))
	if err != nil {
		return "", false
	}

	// len of string = 5 fixed bit + (len(s[1])) + wordLen
	if len(s[0]) != 5+(len(s[1]))+wordLen {
		return "", false
	}

	return s[0][3+(len(s[1])) : 3+(len(s[1]))+wordLen], true
}

func StatusResponse(s string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", s))
}

func NullResponse() []byte {
	return []byte("$-1\r\n")
}

func BulkStringResponse(s string) []byte {
	return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(s), s))
}
