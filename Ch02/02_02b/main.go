package main

import (
	"fmt"
)

func main() {

	str1 := "John has"
	str2 := "not finished"
	str3 := "his homework"

	// Whenever we call print functions, we get back
	// two values - first is a numeric value which is the length of the string
	// and second is the error object.
	// We can receive these values by declaring the variables at the beginning of the statement

	len, err := fmt.Println(str1, str2, str3)
	fmt.Println("Error is", err)
	fmt.Println("Length is", len)
}
