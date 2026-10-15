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
