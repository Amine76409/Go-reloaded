package main

import "testing"

func TestFixArticles(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{

		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"a hour ago", "an hour ago"},
		{"a cat and a dog", "a cat and a dog"},
		{"a", "a"},
		{"a end", "an end"},
	}
}