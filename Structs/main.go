package main

import (
	"fmt"
)

type contactInfo struct {
	number int
	email string
}

type person struct {
	firstname string
	lastname string
	contactInfo
}

func main() {
	foo := person{
		firstname: "John",
		lastname: "Doe",
		contactInfo: contactInfo{
			number: 1234,
			email: "foo@foo.com",
		},
	}

	foo.print()
	foo.updateName("Jane")
	foo.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newName string) {
	(*p).firstname = newName
}