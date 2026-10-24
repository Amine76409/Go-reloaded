package main

import (
	"strings"
	"testing"
)

func TestTokenize(t *testing.T) {
	// "(up, 2)" doit rester en un seul morceau
	got := Tokenize("This is so exciting (up, 2) now")
	// on met un | entre les mots pour pouvoir comparer
	result := strings.Join(got, "|")
	if result != "This|is|so|exciting|(up, 2)|now" {
		t.Errorf("got %q", result)
	}
}

// la virgule après le marqueur doit être un mot à part
func TestTokenizeMarkerWithPunctuation(t *testing.T) {
	got := Tokenize("times (up) , it")
	result := strings.Join(got, "|")
	if result != "times|(up)|,|it" {
		t.Errorf("got %q", result)
	}
}

func TestIsMarker(t *testing.T) {
	// ce sont des marqueurs
	if !IsMarker("(hex)") {
		t.Error("(hex) est un marqueur")
	}
	if !IsMarker("(bin)") {
		t.Error("(bin) est un marqueur")
	}
	if !IsMarker("(up)") {
		t.Error("(up) est un marqueur")
	}
	if !IsMarker("(low)") {
		t.Error("(low) est un marqueur")
	}
	if !IsMarker("(cap)") {
		t.Error("(cap) est un marqueur")
	}
	// avec un nombre
	if !IsMarker("(up, 3)") {
		t.Error("(up, 3) est un marqueur")
	}
	if !IsMarker("(cap, 12)") {
		t.Error("(cap, 12) est un marqueur")
	}

	// ce ne sont pas des marqueurs
	if IsMarker("hello") {
		t.Error("hello n'est pas un marqueur")
	}
	// ce marqueur n'existe pas
	if IsMarker("(hello)") {
		t.Error("(hello) n'est pas un marqueur")
	}
	// la parenthèse n'est pas fermée
	if IsMarker("(up") {
		t.Error("(up n'est pas un marqueur")
	}
}
