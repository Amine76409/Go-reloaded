package main

import (
	"reflect"
	"testing"
)

func TestCapitalize(t *testing.T) {
	tests := map[string]string{
		"bridge": "Bridge",
		"BRIDGE": "Bridge",
		"a":      "A",
		"":       "",
	}
	for word, want := range tests {
		got := Capitalize(word)
		if got != want {
			t.Errorf("Capitalize(%q) = %q, want %q", word, got, want)
		}
	}
}

func TestApplyCase(t *testing.T) {
	got := ApplyCase([]string{"This", "is", "so", "exciting"}, "up", 2)
	want := []string{"This", "is", "SO", "EXCITING"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestApplyCaseTooMany(t *testing.T) {
	got := ApplyCase([]string{"a", "b"}, "cap", 10)
	want := []string{"A", "B"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCaseInText(t *testing.T) {
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
	for _, test := range tests {
		got := Process(test.input)
		if got != test.want {
			t.Errorf("Process(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
