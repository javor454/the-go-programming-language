package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("command name: %s\nargs: %v", os.Args[0], strings.Join(os.Args[1:], " "))
}
