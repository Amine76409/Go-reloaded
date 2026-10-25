package main

import "testing"

func TestFixQuotes(t *testing.T) {
	// les espaces dans les guillemets sont enlevés
	if FixQuotes("I am exactly how they describe me: ' awesome '") != "I am exactly how they describe me: 'awesome'" {
		t.Error("' awesome ' doit devenir 'awesome'")
	}

	// pareil avec plusieurs mots dans les guillemets
	if FixQuotes("As Elton John said: ' I am the most well-known homosexual in the world '") != "As Elton John said: 'I am the most well-known homosexual in the world'" {
		t.Error("les espaces dans la phrase entre guillemets ne sont pas enlevés")
	}

	// le ' de don't est dans un mot, il doit rester
	if FixQuotes("I don't know") != "I don't know" {
		t.Error("don't ne doit pas changer")
	}

	// déjà bien écrit, donc ça ne change pas
	if FixQuotes("'already fine'") != "'already fine'" {
		t.Error("'already fine' ne doit pas changer")
	}
}

func TestQuotesInText(t *testing.T) {
	// on teste avec toutes les règles ensemble (guillemets et virgule)
	got := Process("He said ' hello there ' , then left")
	want := "He said 'hello there', then left"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
