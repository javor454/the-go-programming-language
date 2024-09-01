package main

import (
	"bytes"
	"fmt"
)

func main() {
	// TODO page 45
	// exercise 2.3
	// exercise 2.4
	// exercise 2.5
	numbers := []int{1, 2, 3, 4, 5}

	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range numbers {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	fmt.Println(buf.String())
}
