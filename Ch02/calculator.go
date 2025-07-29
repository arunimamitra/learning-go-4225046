package main

// Uncomment this import section to use required Go packages
import (
	//"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the the result of adding 2 numbers
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	val1, _ := strconv.ParseFloat(strings.TrimSpace(value1), 64)

	// Convert the second string to a float64
	val2, _ := strconv.ParseFloat(strings.TrimSpace(value2), 64)

	// Calculate and return the result
	val3 := val1 + val2

	return val3
}
