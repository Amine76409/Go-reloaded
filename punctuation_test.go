package main

import "testing"

func TestFixPunctuation(t *testing.T) {
	if FixPunctuation("I was sitting over there ,and then BAMM !!") != "I was sitting over there, and then BAMM!!" {
		t.Error("la virgule et les !! ne sont pas bien placés")
	}

	// les trois points restent ensemble
	if FixPunctuation("I was thinking ... You were right") != "I was thinking... You were right" {
		t.Error("les ... ne sont pas bien placés")
	}

	// plusieurs ponctuations dans la même phrase
	if FixPunctuation("Punctuation tests are ... kinda boring ,what do you think ?") != "Punctuation tests are... kinda boring, what do you think?" {
		t.Error("la phrase avec ... , et ? n'est pas bonne")
	}

	// le point-virgule et les deux-points
	if FixPunctuation("hello ; world : yes") != "hello; world: yes" {
		t.Error("; et : ne sont pas bien placés")
	}

	// !? reste ensemble
	if FixPunctuation("really !? no") != "really!? no" {
		t.Error("!? doit rester ensemble")
	}

	// dans un nombre, on ne change rien
	if FixPunctuation("pi is 3.14 and it is 10:30") != "pi is 3.14 and it is 10:30" {
		t.Error("les nombres ne doivent pas changer")
	}

	// ponctuation à la fin : pas d'espace après
	if FixPunctuation("the end .") != "the end." {
		t.Error("le point final est mal placé")
	}
}
