package main

import (
	"fmt"
)

func main() {

	anInt := 42
	var p *int = &anInt

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Value of p:", *p)
	}

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Pointer:", *pointer1)
	fmt.Println("Original pointer addr:", pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("\nPointer:", *pointer1)
	fmt.Println("Original value:", value1)
	// Point is to show that when we change the pointer that's pointing
	// to a variable,
	// we are changing the underlying variable itself.
}
