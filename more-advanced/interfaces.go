package main

import "fmt"

//this means main function for the first example
func main1()  {
	
	var x interface{}

	x = "Hello, World!"
	fmt.Println(x)
	x = 100
	fmt.Println(x)


	s, ok := x.(int)

	if !ok {
		panic("NO!")
	}
	fmt.Println(s)

}

//this means main function for the second example
func main2()  {
	
	var x interface{}

	x = "Hello, World"
	// x = 12
	switch x.(type) {
	case int:
		fmt.Println("Integer!")
	case string:
		fmt.Println("String!")
	}
}


//This area to the end belongs to the 3rd example


type BasePerson struct {
	First	string
	Last	string
}

func (bp BasePerson) GetName() string {
	return bp.First
}

type Employee struct {
	BasePerson
	Salary int
	Boss *Manager
}

type Manager struct {
	Employee
}

type Person interface {
	GetName() string
}

func SayHello(p Person){
	fmt.Printf("Hello, %s\n", p.GetName())
}

func main()  {
	
	m := &Manager{
		Employee{
			BasePerson: BasePerson{
				First:	"John",
				Last:	"Doe",
			},
			Salary: 40000,
			Boss: nil,
		},
	}

	e := &Employee{
		BasePerson: BasePerson{
			First: "Steve",
			Last: "Jones",
		},
		Salary: 30000,
		Boss:	m,
	}

	fmt.Println(m.First)
	fmt.Println(e.First)

	SayHello(m)
	SayHello(e)

	people := []Person{Person(m), Person(e)}
	for _, person := range people {
		SayHello(person)
	}

}