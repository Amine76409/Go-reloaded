package main

import "strings"

func FixQuotes(text string) string {
	runes := []rune(text)
	result := ""
	// open = true si on est dans des guillemets
	open := false
	// skip = true si on doit ignorer les espaces
	skip := false

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// on ignore les espaces juste après un guillemet qui ouvre
		if skip && r == ' ' {
			continue
		}
		skip = false

		// si ce n'est pas un guillemet, on le garde
		if r != '\'' {
			result = result + string(r)
			continue
		}

		// un ' au milieu d'un mot (ex : don't), on le garde
		if i > 0 && i+1 < len(runes) && isLetter(runes[i-1]) && isLetter(runes[i+1]) {
			result = result + string(r)
			continue
		}

		if open == false {
			// guillemet qui ouvre
			result = result + "'"
			open = true
			skip = true
		} else {
			// guillemet qui ferme : on enlève les espaces avant
			result = strings.TrimRight(result, " ")
			result = result + "'"
			open = false
		}
	}
	return result
}
