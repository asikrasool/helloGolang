package main

import (
	"fmt"
)

type contactinfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {
	// One way to print struct
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	//Secoundway to print struct
	// Defining structs with empty value
	var john person
	fmt.Printf("%+v", john)

	//Update values to struct
	john.firstName = "John"
	john.lastName = "Boyle"
	// fmt.Printf("%+v", john)
	john.print()

	//Embeding one struct with another
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactinfo{
			email:   "jim@gmail.com",
			zipCode: 627005,
		},
	}
	jim.updateName("Jimmy")
	jim.print() // This line, still gonna print "Jim" as first line. Need of pointers
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
