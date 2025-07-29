package main

import (
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	var val1 float64 = convertInputToValue(input1)
	var val2 float64 = convertInputToValue(input2)
	switch operation {
	case "+":
		return addValues(val1, val2)
	case "-":
		return subtractValues(val1, val2)
	case "*":
		return multiplyValues(val1, val2)
	case "/":
		return divideValues(val1, val2)
	}
	return 0
}

func convertInputToValue(input string) float64 {
	res, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// if err != nil {
	// 	fmt.Printf("panic: %v must be a number", input)
	// 	return 0
	// }
	// fmt.Printf("value is: %v ", res)
	return res
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
