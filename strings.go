package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I'm  a String"

	fmt.Printf("Ends with string? %v\n", strings.HasSuffix(strings.ToLower(s), "string"))

}