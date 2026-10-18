package main

import "unicode"

func FixQuotes(text string) string {
	runes := []rune(text)
	result := []rune{}
	open := false
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		insideWord := i > 0 && i+1 < len(runes) &&
			unicode.IsLetter(runes[i-1]) && unicode.IsLetter(runes[i+1])
		if r != '\'' || insideWord {
			result = append(result, r)
			continue
		}

		if !open {
			result = append(result, r)
			open = true
			for i+1 < len(runes) && runes[i+1] == ' ' {
				i++
			}
		} else {
			for len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}
			result = append(result, r)
			open = false
		}
	}
	return string(result)
}
