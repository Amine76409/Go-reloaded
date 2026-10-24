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

// vrai si la ponctuation à la place i est dans un nombre (3.14 ou 10:30)
func dansUnNombre(runes []rune, i int) bool {
	if i == 0 || i == len(runes)-1 {
		return false
	}
	return isDigit(runes[i-1]) && isDigit(runes[i+1])
}

func FixPunctuation(text string) string {
	// étape 1 : on enlève l'espace avant chaque ponctuation
	for _, p := range ".,!?:;" {
		text = strings.ReplaceAll(text, " "+string(p), string(p))
	}

	// étape 2 : on ajoute un espace après la ponctuation
	runes := []rune(text)
	result := ""
	for i, r := range runes {
		result = result + string(r)

		// si c'est la fin du texte, on ne met rien
		if i == len(runes)-1 {
			continue
		}
		// on met un espace si la ponctuation est suivie d'une lettre
		// (pas d'espace pour ... ou !? ou un nombre comme 3.14)
		suivant := runes[i+1]
		if isPunct(r) && suivant != ' ' && !isPunct(suivant) && !dansUnNombre(runes, i) {
			result = result + " "
		}
	}
	return result
}
