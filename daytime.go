package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now().Hour()
	current_greet := check_day(current_time)
	fmt.Println(current_greet)
}

func check_day(hour int) string {
	var day_greet string
	switch {
	case hour < 12:
		day_greet = "Good Morning"
	case hour >= 12:
		day_greet = "Good Afternoon"
	default:
		day_greet = "No Idea"
	}
	return day_greet
}
