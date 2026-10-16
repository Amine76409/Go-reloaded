package main

import (
	"strconv"
	"strings"
)

func ReadMarker(token string) (string, int) {
	inside := strings.Trim(token, "()")
	parts := strings.Split(inside, ";")
	name := strings.TrimSpace(parts[0])
	n := 1
	if len(parts) > 1 {
			number, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err == nil {
					n = number
			}
	}
	return name, n
}