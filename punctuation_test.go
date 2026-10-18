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
	
		isNumber := (runes[i] == '.' || runes[i] == ',' || runes[i] == ':') &&
			i > 0 && i+1 < len(runes) &&
			unicode.IsDigit(runes[i-1]) && unicode.IsDigit(runes[i+1])
		if isNumber {
			result = append(result, runes[i])
			i++
			continue
		}

		for len(result) > 0 && result[len(result)-1] == ' ' {
			result = result[:len(result)-1]
		}
		for i < len(runes) && isPunct(runes[i]) {
			result = append(result, runes[i])
			i++
		}
		for i < len(runes) && runes[i] == ' ' {
			i++
		}
		if i < len(runes) {
			result = append(result, ' ')
		}
	}
	return string(result)
}