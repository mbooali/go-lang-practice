package main

import(
	"fmt"
	"strconv"
	"os"
)

func main()  {
	// always initialize the values

	var sum int

	for _, a := range os.Args[1:] {

		i, err := strconv.Atoi(a)

		if err != nil {
			panic(fmt.Sprintf("Invalid Value: %v", err))
		}
		sum += i

	}

	fmt.Printf("Sum = %v\n", sum)
}