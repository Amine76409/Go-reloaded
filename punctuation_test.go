package main

import (
	"strings"
	"unicode"
)

func isPunct(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

func FixPunctuation(text string) string {
	runes := []rune(text)
	result := []rune{}
	i := 0
	for i < len(runes) {
		if !isPunct(runes[i]) {
			result = append(result, runes[i])
			i++
			continue
		}
	}
}