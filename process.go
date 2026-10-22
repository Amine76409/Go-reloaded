package main

import (
	"strconv"
	"strings"
)

func ReadMarker(token string) (string, int) {
	// on enlève les parenthèses
	inside := strings.Trim(token, "()")
	// on sépare avec la virgule
	parts := strings.Split(inside, ",")
	name := strings.TrimSpace(parts[0])

	// n = le nombre de mots à changer (1 par défaut)
	n := 1
	// s'il y a un nombre après la virgule, on le prend
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
		// si ce n'est pas un marqueur, on garde le mot
		if !IsMarker(token) {
			result = append(result, token)
			continue
		}
		// s'il n'y a pas de mot avant, on ne fait rien
		if len(result) == 0 {
			continue
		}

		name, n := ReadMarker(token)
		// last = la place du dernier mot
		last := len(result) - 1

		if name == "hex" {
			result[last] = ConvertNumber(result[last], 16)
		} else if name == "bin" {
			result[last] = ConvertNumber(result[last], 2)
		} else {
			// up, low ou cap
			result = ApplyCase(result, name, n)
		}
	}
	return result
}

func ProcessLine(line string) string {
	// on coupe la ligne en mots
	tokens := Tokenize(line)
	// on applique les marqueurs
	tokens = ApplyMarkers(tokens)
	// on remet les mots ensemble
	text := strings.Join(tokens, " ")
	// on corrige la ponctuation, les guillemets et les "a"
	text = FixPunctuation(text)
	text = FixQuotes(text)
	text = FixArticles(text)
	return text
}

func Process(text string) string {
	// on coupe le texte en lignes
	lines := strings.Split(text, "\n")
	// on corrige chaque ligne
	for i, line := range lines {
		lines[i] = ProcessLine(line)
	}
	return strings.Join(lines, "\n")
}
