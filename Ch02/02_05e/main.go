package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("The value of PI is:", math.Pi)

	circleRadius := 15.0
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.02f\n", circumference)
}
