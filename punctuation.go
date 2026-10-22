package main

import "strings"

// vrai si le caractère est de la ponctuation
func isPunct(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

// vrai si le caractère est un chiffre
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func FixPunctuation(text string) string {
	runes := []rune(text)
	result := ""
	// needSpace = true si on doit mettre un espace après la ponctuation
	needSpace := false

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// après la ponctuation, on ignore les espaces
		if needSpace && r == ' ' {
			continue
		}

		// si ce n'est pas de la ponctuation
		if !isPunct(r) {
			// on met l'espace qu'on avait gardé
			if needSpace {
				result = result + " "
				needSpace = false
			}
			result = result + string(r)
			continue
		}

		// si c'est dans un nombre (3.14 ou 10:30), on ne change rien
		if (r == '.' || r == ',' || r == ':') && i > 0 && i+1 < len(runes) && isDigit(runes[i-1]) && isDigit(runes[i+1]) {
			result = result + string(r)
			continue
		}

		// ponctuation : on enlève les espaces avant et on l'ajoute
		result = strings.TrimRight(result, " ")
		result = result + string(r)
		needSpace = true
	}
	return result
}
