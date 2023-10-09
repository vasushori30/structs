package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName   string
	lastName    string
	contactInfo // this is a struct inside a struct
}

func main() {
	james := person{
		firstName: "James",
		lastName:  "Bond",
		contactInfo: contactInfo{
			email:   "james.bond007@gmail.com",
			zipCode: 007,
		},
	}

	jamesPointer := &james // This gives memory address of the value this variable is pointing at
	jamesPointer.updateName("Jimmy")
	james.print()
}

// Turn value into address with &value
// Turn address into value with *address

// * person is a type description, it means we are working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// Here *pointerToPerson gives the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
