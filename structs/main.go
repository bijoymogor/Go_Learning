package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// func main() {
// 	// alex := person{"Alex","Anderson"}

// 	alex := person{
// 		firstName: "Alex",
// 		lastName:  "Anderson"}

// 	fmt.Println(alex)

// 	var bijoy person
// 	fmt.Println(bijoy)

// 	fmt.Printf("%+v", bijoy)

// 	bijoy.firstName = "Bijoy"
// 	bijoy.lastName = "Mogor"

// 	fmt.Printf("%+v", bijoy)

// }

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 100,
		},
	}

	// Call by value
	jim.updateName("Nero")

	// jimPointer := &jim
	// jimPointer.updateName("NERO")
	jim.updateName("NERO")
	jim.print()
}

//Call by value
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName

// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)

}
