package main

import "testing"

func TestConvertNumber(t *testing.T) {
	// chaque test : un mot, une base, et ce qu'on veut à la fin
	tests := []struct {
		word string
		base int
		want string
	}{
		{"1E", 16, "30"},
		{"42", 16, "66"},
		{"10", 2, "2"},
		{"1010", 2, "10"},
		{"zz", 16, "zz"},// pas un nombre, donc ça ne change pas
		{"12", 2, "12"},// 2 n'existe pas en binaire, donc ça ne change pas
	}

	// on teste chaque cas un par un
	for _, test := range tests {
		got := ConvertNumber(test.word, test.base)
		// si le résultat n'est pas bon, on affiche l'erreur
		if got != test.want {
			t.Errorf("ConvertNumber(%q, %d) = %q, want %q", test.word, test.base, got, test.want)
		}
	}
}

func TestHexAndBinInText(t *testing.T) {
	// on teste (hex) et (bin) dans une phrase
	got := Process("Simply add 42 (hex) and 10 (bin) and you will see the result is 68.")
	want := "Simply add 66 and 2 and you will see the result is 68."
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
