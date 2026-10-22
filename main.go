package main

import (
	"fmt"
	"os"
)

func main() {
	// il faut 2 arguments : le fichier à lire et le fichier à écrire
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	// on lit le fichier
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// on corrige le texte
	result := Process(string(data))

	// on écrit le résultat dans l'autre fichier
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
