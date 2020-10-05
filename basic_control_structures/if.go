package main

import (
	"fmt"
)

func main() {

	a := 5

	if a < 0 {
		fmt.Println("Negative Value")
	} else if a < 10 {
		fmt.Println("Single Digit")
	} else {
		fmt.Println("Multiple Digits")
	}
}