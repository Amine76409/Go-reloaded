package main

import "testing"

func TestConvertNumber(t *testing.T) {
	tests := []struct {
		word string
		base int
		want string
	}{
		{"1E", 16, "30"},
		{"42", 16, "66"},
		{"10", 2, "2"},
		{"1010", 2, "10"},
		{"zz", 16, "zz"},
		{"12", 2, "12"},
	}
	for _, test := range tests {
		got := ConvertNumber(test.word, test.base)
		if got != test.want {
			t.Errorf("ConvertNumber(%q, %d) = %q, want %q", test.word, test.base, got, test.want)
		}
	}
}

func TestHexAndBinInText(t *testing.T) {
	got := Process("Simply add 42 (hex) and 10 (bin) and you will see the result is 68.")
	want := "Simply add 66 and 2 and you will see the result is 68."
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
