package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	count := len(msg)
	new_string := strings.Repeat("!", count) + strings.ToUpper(msg) + strings.Repeat("!", count)
	fmt.Println(new_string)
}
