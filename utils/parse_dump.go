package utils

import "strings"

func GetDumpTable(s string) []string {
	lines := strings.Split(s, "\n")

	for i := range lines {
		lines[i] = lines[i][10:59]
	}

	return lines
}
