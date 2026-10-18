package main

import "testing"

func TestFixPunctuation(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},
		{"hello ; world : yes", "hello; world: yes"},
		{"really !? no", "really!? no"},
		{"pi is 3.14 and it is 10:30", "pi is 3.14 and it is 10:30"},
		{"the end .", "the end."},
	}
	for _, test := range tests {
		got := FixPunctuation(test.input)
		if got != test.want {
			t.Errorf("FixPunctuation(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
