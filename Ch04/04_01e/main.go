package main

import (
	"fmt"
)

func main() {
	// theAnswer := 0
	var result string

	// Very similar to java and C, only we don't need paranthesis in the conditional statements
	// And we have the option to define variable inside conditional block like line 14

	if theAnswer := 42; theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)
}
