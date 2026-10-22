package main

import "strings"

func FixArticles(text string) string {
	// on sépare le texte en mots
	words := strings.Split(text, " ")

	// on s'arrête avant le dernier mot, car il n'y a rien après
	for i := 0; i < len(words)-1; i++ {
		// si le mot n'est pas "a" ou "A", on passe
		if words[i] != "a" && words[i] != "A" {
			continue
		}

		// on enlève l'apostrophe au début du mot suivant
		next := strings.TrimLeft(words[i+1], "'")
		if next == "" {
			continue
		}

		// première lettre du mot suivant, en minuscule
		first := strings.ToLower(next[:1])

		// si c'est une voyelle ou un h, on ajoute "n"
		if strings.Contains("aeiouh", first) {
			words[i] = words[i] + "n"
		}
	}
	return strings.Join(words, " ")
}
