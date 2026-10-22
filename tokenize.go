package main

import "strings"

func Tokenize(line string) []string {
	// on coupe aux espaces
	words := strings.Fields(line)
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]
		// "(up," et "2)" vont ensemble : on les colle
		if strings.HasSuffix(word, ",") && i+1 < len(words) {
			word = word + " " + words[i+1]
			if IsMarker(word) {
				i++
			} else {
				word = words[i]
			}
		}
		result = append(result, word)
	}
	return result
}

func IsMarker(token string) bool {
	// les marqueurs possibles
	markers := []string{"(hex)", "(bin)", "(up)", "(low)", "(cap)"}
	for _, m := range markers {
		if token == m {
			return true
		}
	}
	// avec un nombre : (up, 2)
	if strings.HasPrefix(token, "(up, ") || strings.HasPrefix(token, "(low, ") || strings.HasPrefix(token, "(cap, ") {
		return strings.HasSuffix(token, ")")
	}
	return false
}
