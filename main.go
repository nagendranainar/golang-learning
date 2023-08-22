package main

import "fmt"

func main() {
	a := 10
	var b string
	card := newCard()
	fmt.Println(card)
	fmt.Println(a)
	fmt.Println(b)
}

func newCard() string {
	return "Cards"
}
