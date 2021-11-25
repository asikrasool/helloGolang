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
	name := "Bill"

	fmt.Println(*&name)

	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	//Another way to print struct
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
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	// The above pointer can also be implemented by just type of var instead of pointer of vat Ex:
	jim.updateName("Jimmy") // this call update name fuction without pointer of jim var, since updateName function expects pointer type of person(jim)
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// changing below code to make use of pointers
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
