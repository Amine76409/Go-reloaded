package main

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, it"},
		{"", ""},
		{"zz (hex) word", "zz word"},
		{"(up) hello", "hello"},
	}

	for _, test := range tests {
		got := Process(test.input)
		if got != test.want {
			t.Errorf("Process(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestProcessManyLines(t *testing.T) {
	got := Process("hello (up)\nworld (cap)\n")
	want := "HELLO\nWorld\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
