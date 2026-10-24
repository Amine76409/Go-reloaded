package main

import "testing"

func TestProcess(t *testing.T) {
	// chaque test : une phrase de départ et ce qu'on veut à la fin
	tests := []struct {
		input string
		want  string
	}{
		// (hex) et (bin) : on change le nombre en décimal
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},

		// (up), (low) et (cap) : on change les majuscules
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},

		// la ponctuation
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},

		// les guillemets
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},

		// "a" devient "an" devant une voyelle
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},

		// une phrase avec plusieurs règles en même temps
		{"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, it"},

		// cas limites
		{"", ""},                     // texte vide
		{"zz (hex) word", "zz word"}, // zz n'est pas un nombre, on enlève juste le marqueur
		{"(up) hello", "hello"},      // pas de mot avant le marqueur
	}

	// on teste chaque phrase une par une
	for _, test := range tests {
		got := Process(test.input)
		// si le résultat n'est pas bon, on affiche l'erreur
		if got != test.want {
			t.Errorf("Process(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

// on vérifie que ça marche avec plusieurs lignes
func TestProcessManyLines(t *testing.T) {
	got := Process("hello (up)\nworld (cap)\n")
	want := "HELLO\nWorld\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
