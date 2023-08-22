package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	name := strings.ToUpper(os.Args[1])

	fmt.Printf("hmm....so your name is %s", name)
}
