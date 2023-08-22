//This example is to create different struct data structure.

package main

import "fmt"

//Below is a struct that groups all relevant
//data of a person.

type person struct {
	FirstName string
	LastName  string
	Age       int
	contact   contact
}

//Below is a struct that is embedded within other structs

type contact struct {
	email   string
	zipcode int
}

func main() {

	boss := person{
		FirstName: "Lavanya",
		LastName:  "Rajaram Shanthi",
		Age:       38,
		contact: contact{
			email:   "lavydingi@gmail.com",
			zipcode: 27709,
		},
	}

	driver := person{
		FirstName: "Nagendra Kumar",
		LastName:  "Nainar",
		Age:       43,
		contact: contact{
			email:   "nagudingan@gmail.com",
			zipcode: 27713,
		},
	}

	fmt.Println(boss.FirstName)
	fmt.Printf("%+v", boss)
	fmt.Println(driver)

}
