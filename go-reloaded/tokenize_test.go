package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	got := Tokenize("This is so exciting (up, 2) now")
	want := []string{"This", "is", "so", "exciting", "(up, 2)", "now"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTokenizeMarkerWithPunctuation(t *testing.T) {
	got := Tokenize("times (up) , it")
	want := []string{"times", "(up)", ",", "it"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIsMarker(t *testing.T) {
	tests := map[string]bool{
		"(hex)":     true,
		"(bin)":     true,
		"(up)":      true,
		"(low)":     true,
		"(cap)":     true,
		"(up, 3)":   true,
		"(cap, 12)": true,
		"hello":     false,
		"(hello)":   false,
		"(up":       false,
	}
	for token, want := range tests {
		got := IsMarker(token)
		if got != want {
			t.Errorf("IsMarker(%q) = %v, want %v", token, got, want)
		}
	}
}
