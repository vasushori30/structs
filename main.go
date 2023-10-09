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

	james.updateName("Jimmy")
	james.print()
}

// Type of *person or a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// Here *pointerToPerson gives the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
