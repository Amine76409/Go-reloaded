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

func ApplyMarkers(tokens []string) []string {
	result := []string{}
	for _, token := range tokens {
		if !IsMarker(token) {
			result = append(result, token)
			continue
		}
		if len(result) == 0 {
			continue
		}
		name, n := ReadMarker(token)
		last := len(result) - 1
		if name == "hex" {
			result[last] = ConvertNumber(result[last], 16)
		} else if name == "bin" {
			result[last] = ConvertNumber(result[last], 2)
		} else {
			result = ApplyCase(result, name, n)
		}
	}
	return result
}

func ProcessLine(line string) string {
	tokens := Tokenize(line)
	tokens = ApplyMarkers(tokens)
	text := strings.Join(tokens, " ")
	text = FixPunctuation(text)
	text = FixQuotes(text)
	text = FixArticles(text)
	return text
}

func Process(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = ProcessLine(line)
	}
	return strings.Join(lines, "\n")
}
