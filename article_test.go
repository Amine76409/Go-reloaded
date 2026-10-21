package main

import "testing"

func TestFixArticles(t *testing.T) {
	// la liste des tests : le texte de depart et ce qu'on veut a la fin
	tests := []struct {
		input string
		want  string
	}{
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"a hour ago", "an hour ago"},   // le h compte aussi
		{"a cat and a dog", "a cat and a dog"}, // consonne donc ca change pas
		{"a", "a"},     // y'a pas de mot apres donc ca doit pas planter
		{"a end", "an end"},
	}

	// on teste chaque cas un par un
	for _, test := range tests {
		got := FixArticles(test.input)
		// si c'est pas pareil on affiche l'erreur
		if got != test.want {
			t.Errorf("FixArticles(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
