package main

import "fmt"

func main()  {
	m1 := map[string]string{}

	m1["first"] = "John"
	m1["last"] = "Doe"

	fmt.Println(m1)

	m2 := map[string]string{
		"first": "John",
		"last": "Doe",
	}

	fmt.Println(m2)

	m3 := make(map[string]string)

	m3["first"] = "John"
	m3["last"] = "Doe"
		
	fmt.Println(m3["first"])
}