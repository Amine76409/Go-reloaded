package main

import (
	"strings"
	"unicode"
)

func FixArticles(text string) string {
	words := strings.Split(text, " ")
	for i := 0; i < len(words)-1; i++ {
		if words[i] != "a" && words[i] != "A" {
			continue
		}
		next := strings.TrimLeft(words[i+1], "'")
		if next == "" {
			continue
		}
		first := unicode.ToLower([]rune(next)[0])
		if strings.ContainsRune("aeiouh", first) {
			words[i] = words[i] + "n"
		}
	}
	return strings.Join(words, " ")
}
