package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a string
	fmt.Fscan(reader, &a)
	var b string
	fmt.Fscan(reader, &b)

	switch {
	case len(a) > len(b):
		fmt.Fprintf(writer, "%v\n", "GREATER")
	case len(a) < len(b):
		fmt.Fprintf(writer, "%v\n", "LESS")
	default:
		if a < b {
			fmt.Fprintf(writer, "%v\n", "LESS")
		} else if b < a {
			fmt.Fprintf(writer, "%v\n", "GREATER")
		} else {
			fmt.Fprintf(writer, "%v\n", "EQUAL")
		}
	}

}

