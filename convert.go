package main

import "strconv"

func ConvertNumber(word string, base int) string {
	// on transforme le mot en nombre (base 16 ou base 2)
	number, err := strconv.ParseInt(word, base, 64)
	// si ce n'est pas un vrai nombre, on rend le mot sans le changer
	if err != nil {
		return word
	}
	// on rend le nombre en décimal, sous forme de texte
	return strconv.Itoa(int(number))
}
