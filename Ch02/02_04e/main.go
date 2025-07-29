package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	// prints 164.8999999 - precision problem, common w other languages too
	// Float values are tracked internally using bits(binary format)
	fmt.Println("Float sum:", floatSum)

	// In java we can solve this using BigDecimal API
	// Big decimal in Go is available but it will result in the same answer
	// instead we can use "round" function from math package
	roundV := math.Round(floatSum)
	fmt.Println("Round sum:", roundV) // this will round up to nearest integer

	roundV = math.Round(floatSum*100) / 100 // this will eliminate some of the precision
	fmt.Println("Round sum2:", roundV)

	// Wrapping integer inside a float64
	total := float64(i1) + f2
	fmt.Println("Result:", total)

	product := float64(i1) * f2
	fmt.Println("Product:", product)

	remainder := i2 % i1
	fmt.Println(remainder)
}
