package main

import(
	"fmt"
)


func main()  {
	s := []int{}
	fmt.Println(s)

	s = []int{1, 2, 3, 4}
	fmt.Println(s)

	s = append(s, 5)

	s = make([]int, 0)
	fmt.Println(s)

	s = make([]int, 10)
	fmt.Println(s)

	s[2] = -999
	fmt.Println(s)


	r := [4]int{1, 2, 3}
	fmt.Println(r)

}