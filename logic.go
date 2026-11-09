package main

import "strings"

// Greeting returns a friendly message for the supplied name.
func Greeting(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "developer"
	}
	return "Hello, " + name + "!"
}

// SumRange returns the inclusive sum of all integers from from to.
// A reversed range contains no values and returns zero.
func SumRange(from, to int) int {
	if from > to {
		return 0
	}

	sum := 0
	for value := from; value <= to; value++ {
		sum += value
	}
	return sum
}
