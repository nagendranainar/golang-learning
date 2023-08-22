package main

import (
	"fmt"
	"os"
)

var (
	credential_list = make(map[string]string)
)

func main() {

	credential_list["Nagendra"] = "cisco"
	credential_list["Lavanya"] = "ceo"
	//uname := []string{"Nagendra", "Lavanya"}
	//passwd := []string{"cisco", "ceo"}

	if len(os.Args) < 3 {
		fmt.Println("Please enter both the username and password")
	} else {
		passwd, uname_exist := credential_list[os.Args[1]]
		if uname_exist {
			if passwd == os.Args[2] {
				fmt.Println("Access granted")
			} else {
				fmt.Println("Access Denied")
			}
		} else {
			fmt.Println("Username %d Not Found \n", os.Args[1])
		}
	}

	//if len(os.Args) < 3 {
	//	fmt.Println("Please enter both username and password")
	//} else if slices.Contains(uname, os.Args[1]) && slices.Contains(passwd, os.Args[2]) {
	//	fmt.Println("You are Allowed")
	//} else {
	//	fmt.Println("Access Denied")
	//}
}
