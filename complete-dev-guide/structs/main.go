package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	// contact contactInfo (if we wanted a specific name for the field)
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (p *person) updateFirstName(value string) {
	(*p).firstName = value
}

func main() {
	// // kinda crappy
	// alex := person{"Alex", "Anderson", contactInfo{"aa@gmail.com", 123456}}

	// // most common
	// judy := person{firstName: "Judy", lastName: "Anderson", contactInfo: contactInfo{email: "ja@gmail.com", zip: 234567}}

	// not that common
	var max person
	var maxContact contactInfo

	maxContact.email = "mm@gmail.com"
	maxContact.zip = 99999

	max.firstName = "Max"
	max.lastName = "Maxxy"

	max.contactInfo = maxContact

	// alex.print()
	// judy.print()

	max.print()

	max.updateFirstName("Maxxy")

	(&max).print()
}
