package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid Argument. Please enter the feet to calculate")
	} else {
		feet, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println("Some error in conversion")
			return
		}

		meter := float64(feet) * 0.3048
		fmt.Println("Output: %d", meter)
	}
}
