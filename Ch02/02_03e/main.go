package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// os.Stdin reads from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	// In Go if we want to ignore an error, we use _ for it's name
	// ReadString - function of the reader object which will read from console
	// returns two values - string itself and the error
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	// only use colon when declaring the variable (line 18)
	// not required in line 24. only = will work
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}
