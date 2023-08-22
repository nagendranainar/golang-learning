package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Enter the number")
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	out := sum_num(num)

	fmt.Println(out)

}

func sum_num(input int) int {
	var sum int
	for i := 0; i <= input; i++ {
		sum += i
	}
	return sum
}
