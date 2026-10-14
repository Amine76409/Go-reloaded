package main

import (
	"strings"
	"unicode"
)

func isPunct(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

