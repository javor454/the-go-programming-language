package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// convert meters to feet
	const factor = 3.28084 // 1m = 3.28084feet

	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments passed to convertor.\nUsage: Convert meters to feet.\nSyntax:$ %s <meters>\n", os.Args[0])
		return
	}

	num := os.Args[1]
	numFloat, err := strconv.ParseFloat(num, 64)
	if err != nil {
		fmt.Printf("Error: %v\nUsage: Convert meters to feet\nSyntax:$ %s <meters>\n", err, os.Args[0])
		return
	}

	feet := numFloat * factor
	fmt.Printf("%f meters are %.2f feet\n", numFloat, feet)
}
