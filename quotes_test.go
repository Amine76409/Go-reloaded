package main

import "testing"

func TestFixQuotes(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"I don't know", "I don't know"},
		{"'already fine'", "'already fine'"},
	}
		for _, test := range tests {
		got := FixQuotes(test.input)
		if got != test.want {
			t.Errorf("FixQuotes(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestQuotesInText(t *testing.T) {
	got := Process("He said ' hello there ' , then left")
	want := "He said 'hello there', then left"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
