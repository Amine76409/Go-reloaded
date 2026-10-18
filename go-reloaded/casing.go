package main

import (
	"strings"
	"unicode"
)

func Capitalize(word string) string {
	runes := []rune(strings.ToLower(word))
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
	}
	return string(runes)
}

func ChangeCase(word string, name string) string {
	if name == "up" {
		return strings.ToUpper(word)
	}
	if name == "low" {
		return strings.ToLower(word)
	}
	return Capitalize(word)
}

func ApplyCase(words []string, name string, n int) []string {
	start := len(words) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = ChangeCase(words[i], name)
	}
	return words
}
