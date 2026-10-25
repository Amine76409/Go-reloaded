package main

import (
	"os"
	"testing"
)

func TestSampleFile(t *testing.T) {
	// on lit le fichier de départ
	input, err := os.ReadFile("sample.txt")
	if err != nil {
		// si on ne peut pas le lire, on arrête le test
		t.Fatal(err)
	}

	// on lit le fichier avec le résultat attendu
	expected, err := os.ReadFile("expected.txt")
	if err != nil {
		t.Fatal(err)
	}

	// on corrige le texte avec notre programme
	got := Process(string(input))

	// si ce n'est pas pareil que le résultat attendu, on affiche l'erreur
	if got != string(expected) {
		t.Errorf("got:\n%s\nwant:\n%s", got, string(expected))
	}
}

