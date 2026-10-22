package main

import (
	"strings"
	"testing"
)

func TestCapitalize(t *testing.T) {
	// "bridge" doit devenir "Bridge"
	if Capitalize("bridge") != "Bridge" {
		t.Error("bridge doit devenir Bridge")
	}
	// même chose si le mot est en majuscules
	if Capitalize("BRIDGE") != "Bridge" {
		t.Error("BRIDGE doit devenir Bridge")
	}
	// un mot d'une seule lettre
	if Capitalize("a") != "A" {
		t.Error("a doit devenir A")
	}
	// un mot vide reste vide
	if Capitalize("") != "" {
		t.Error("un mot vide doit rester vide")
	}
}

func TestApplyCase(t *testing.T) {
	// on met les 2 derniers mots en majuscules
	got := ApplyCase([]string{"This", "is", "so", "exciting"}, "up", 2)
	// on remet les mots en une phrase pour pouvoir comparer
	result := strings.Join(got, " ")
	if result != "This is SO EXCITING" {
		t.Errorf("got %q", result)
	}
}

// si on demande plus de mots qu'il y en a, ça ne doit pas planter
func TestApplyCaseTooMany(t *testing.T) {
	// on demande 10 mots mais il n'y en a que 2
	got := ApplyCase([]string{"a", "b"}, "cap", 10)
	result := strings.Join(got, " ")
	if result != "A B" {
		t.Errorf("got %q", result)
	}
}

func TestCaseInText(t *testing.T) {
	// chaque test : une phrase de départ et ce qu'on veut à la fin
	tests := []struct {
		input string
		want  string
	}{
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"LOUD WORDS HERE (low, 2)", "LOUD words here"},
		{"one two three (cap, 3)", "One Two Three"},
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
