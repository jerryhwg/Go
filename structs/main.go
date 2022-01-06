package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "you@email.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

// original
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// p person: receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
