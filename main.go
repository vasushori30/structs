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
	james.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
