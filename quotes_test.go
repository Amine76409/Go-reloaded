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
}	