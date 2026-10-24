package main

import "testing"

func TestConvertNumber(t *testing.T) {
	// 1E en hexa = 30
	if ConvertNumber("1E", 16) != "30" {
		t.Error("1E en hexa doit donner 30")
	}

	// 42 en hexa = 66
	if ConvertNumber("42", 16) != "66" {
		t.Error("42 en hexa doit donner 66")
	}

	// 10 en binaire = 2
	if ConvertNumber("10", 2) != "2" {
		t.Error("10 en binaire doit donner 2")
	}

	// 1010 en binaire = 10
	if ConvertNumber("1010", 2) != "10" {
		t.Error("1010 en binaire doit donner 10")
	}

	// zz n'est pas un nombre hexa, donc ça ne change pas
	if ConvertNumber("zz", 16) != "zz" {
		t.Error("zz doit rester zz")
	}

	// 2 n'existe pas en binaire, donc ça ne change pas
	if ConvertNumber("12", 2) != "12" {
		t.Error("12 en binaire doit rester 12")
	}
}

func TestHexAndBinInText(t *testing.T) {
	// on teste (hex) et (bin) dans une phrase
	got := Process("Simply add 42 (hex) and 10 (bin) and you will see the result is 68.")
	want := "Simply add 66 and 2 and you will see the result is 68."
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
