package utils

import "strings"

func GetDumpTable(s string) []string {
	return strings.Split(s, "\n")
}
