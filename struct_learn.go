//This example is to create different struct data structure.

package main

import "fmt"

//Below is a struct that groups all relevant
//data of a person.

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	boss := person{
		FirstName: "Lavanya",
		LastName:  "Rajaram Shanthi",
		Age:       38,
	}

	driver := person{
		FirstName: "Nagendra Kumar",
		LastName:  "Nainar",
		Age:       43,
	}

	fmt.Println(boss.FirstName)
	fmt.Printf("%+v", boss)
	fmt.Println(driver)

}
