package main

import(
	"fmt"
	"time"
)

func sayHello(name string)  {
	fmt.Printf("Hello, %s\n", name)
}

func getHello(name string) string{
	return fmt.Sprintf("Hello, %s\n", name)
}

type Person struct {
	First	string
	Last	string
	Dob		time.Time
}

func NewPerson(first, last string, year, month, day int) *Person {
	return &Person{
		First: first,
		Last: last,
		Dob: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}

func (p Person) GetAge() int {
	d := time.Since(p.Dob)
	return int(d.Hours() / 24 / 365)
}

func (p Person) SayHello() {
	fmt.Printf("Hello, %s\n", p.First)
}

func main() {
	sayHello("John")

	s := getHello("John")
	fmt.Println(s)

	p := &Person{
		First:	"John",
		Last:	"Doe",
	}

	p.SayHello()

	new_person := NewPerson("Jane", "Doe", 1980, 1, 1)
	new_person.SayHello()
	fmt.Println(new_person.GetAge())



}