package main

import "fmt"

type Person struct {
	First 	string
	Last 	string
	Age 	int
	Phone	Phone
}

type Phone struct{
	AreaCode	string
	Prefix		string
	Suffix		string
}
/*
func main() {
	p := Person {
		First:	"John",
		Last:	"Doe",
		Age:	30,
		Phone: Phone{
			AreaCode:	"123",
			Prefix:		"354",
			Suffix:		"0123",
		},
	}

	q := &Person{"Jane", "Doe", 25}

	fmt.Println(p)
	fmt.Println(q)
}
*/

func main()  {
	pt := struct{
		X int
		Y int
	} {
		X: 10,
		Y: 20,
	}

	fmt.Println(pt)
}