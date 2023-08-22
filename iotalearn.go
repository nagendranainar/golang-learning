package main

import (
	"fmt"
)

func main() {
	const (
		a = (iota + 1) * 2
		b
		c
	)
	fmt.Println(a, b, c)
}
