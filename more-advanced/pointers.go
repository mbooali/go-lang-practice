package main

import "fmt"

func main() {
	test1()
	test2()
	test3()
	test4()
}

func test1()  {
	s := "My String"

	sptr := &s

	fmt.Println(s)
	fmt.Println(sptr)
	fmt.Println(*sptr)
}

func test2()  {
	sptr := new(string)

	fmt.Println(sptr)
	fmt.Println(*sptr)
}

func test3()  {
	
	s := "My String"
	var sptr *string

	sptr = &s
	fmt.Println(sptr)
}

func test4()  {
	i := new(int)
	fmt.Println(*i)
}