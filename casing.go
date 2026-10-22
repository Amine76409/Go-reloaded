package main

import "strings"

// vrai si le caractère est une lettre (a-z ou A-Z)
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func Capitalize(word string) string {
	// on met tout en minuscule et on transforme en liste de caractères
	runes := []rune(strings.ToLower(word))
	for i, r := range runes {
		// on cherche la première lettre
		if isLetter(r) {
			// on la met en majuscule et on s'arrête
			runes[i] = r - 32
			break
		}
	}
	return string(runes)
}

func ChangeCase(word string, name string) string {
	// up = majuscules
	if name == "up" {
		return strings.ToUpper(word)
	}
	// low = minuscules
	if name == "low" {
		return strings.ToLower(word)
	}
	// sinon c'est cap
	return Capitalize(word)
}

func ApplyCase(words []string, name string, n int) []string {
	// on commence n mots avant la fin
	start := len(words) - n
	// si on demande trop de mots, on commence au début
	if start < 0 {
		start = 0
	}
	// on change les derniers mots
	for i := start; i < len(words); i++ {
		words[i] = ChangeCase(words[i], name)
	}
	return words
}
