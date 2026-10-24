package main

import (
	"strconv"
	"strings"
)

func Process(text string) string {
	// on coupe le texte en lignes
	lines := strings.Split(text, "\n")
	result := ""

	for i, line := range lines {
		words := Tokenize(line)
		newWords := []string{}

		for _, word := range words {
			// si ce n'est pas un marqueur, on garde le mot
			if !IsMarker(word) {
				newWords = append(newWords, word)
				continue
			}
			// s'il n'y a pas de mot avant, on ne fait rien
			if len(newWords) == 0 {
				continue
			}

			// on enlève les parenthèses
			inside := strings.Trim(word, "()")
			// on sépare avec la virgule
			parts := strings.Split(inside, ",")
			name := strings.TrimSpace(parts[0])

			// n = le nombre de mots à changer (1 par défaut)
			n := 1
			if len(parts) > 1 {
				number, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					n = number
				}
			}

			if name == "hex" {
				newWords[len(newWords)-1] = ConvertNumber(newWords[len(newWords)-1], 16)
			} else if name == "bin" {
				newWords[len(newWords)-1] = ConvertNumber(newWords[len(newWords)-1], 2)
			} else {
				// up, low ou cap
				newWords = ApplyCase(newWords, name, n)
			}
		}

		// on remet les mots ensemble
		line = strings.Join(newWords, " ")
		// on corrige la ponctuation, les guillemets et les "a"
		line = FixPunctuation(line)
		line = FixQuotes(line)
		line = FixArticles(line)

		result = result + line
		// on ajoute un retour à la ligne (sauf à la fin)
		if i != len(lines)-1 {
			result = result + "\n"
		}
	}
	return result
}
