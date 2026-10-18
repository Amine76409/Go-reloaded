package main

import (
	"os"
	"testing"
)

func TestSampleFile(t *testing.T) {
	input, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := os.ReadFile("expected.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Process(string(input))
	if got != string(expected) {
		t.Errorf("got:\n%s\nwant:\n%s", got, string(expected))
	}
}
