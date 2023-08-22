package main

import (
	"fmt"
)

func main() {
	var str1 string
	var str2 string

	str1 = "how are you?\n"
	str2 = `how are you`
	fmt.Println(len(str1))
	fmt.Println(str2)
}
