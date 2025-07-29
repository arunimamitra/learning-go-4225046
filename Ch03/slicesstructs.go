package main

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// slicesToObjects() returns a slice of Color objects
func slicesToObjects(colorNames []string, hexValues []int) []Color {
	// Your code goes here.
	// Create a slice of Color objects
	res := make([]Color, 0)
	var c Color
	i := 0
	for k := range colorNames {
		c = Color{colorNames[k], hexValues[k]}
		res = append(res, c)
		i++
	}
	return res
}

type Color struct {
	Name string
	Hex  int
}