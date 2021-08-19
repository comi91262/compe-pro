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
	var s string
	fmt.Fscan(reader, &s)
	var t string
	fmt.Fscan(reader, &t)

	switch s {
	case "Y":
		switch t {
		case "a":
			fmt.Fprintf(writer, "%v\n", "A")
		case "b":
			fmt.Fprintf(writer, "%v\n", "B")
		case "c":
			fmt.Fprintf(writer, "%v\n", "C")
		}

	case "N":
		fmt.Fprintf(writer, "%v\n", t)
	}
}
