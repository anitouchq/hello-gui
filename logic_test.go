package main

import "testing"

func TestGreetingUsesProvidedName(t *testing.T) {
	if got, want := Greeting("Fyne"), "Hello, Fyne!"; got != want {
		t.Fatalf("Greeting() = %q, want %q", got, want)
	}
}

func TestGreetingUsesDefaultForBlankName(t *testing.T) {
	if got, want := Greeting("   "), "Hello, developer!"; got != want {
		t.Fatalf("Greeting() = %q, want %q", got, want)
	}
}

func TestSumRangeIsInclusive(t *testing.T) {
	if got, want := SumRange(1, 10), 55; got != want {
		t.Fatalf("SumRange(1, 10) = %d, want %d", got, want)
	}
}

func TestSumRangeReturnsZeroForReversedRange(t *testing.T) {
	if got := SumRange(10, 1); got != 0 {
		t.Fatalf("SumRange(10, 1) = %d, want 0", got)
	}
}
