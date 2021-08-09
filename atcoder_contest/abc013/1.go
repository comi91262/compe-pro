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

	switch s {
	case "A":
		fmt.Fprintf(writer, "%v\n", 1)
	case "B":
		fmt.Fprintf(writer, "%v\n", 2)
	case "C":
		fmt.Fprintf(writer, "%v\n", 3)
	case "D":
		fmt.Fprintf(writer, "%v\n", 4)
	case "E":
		fmt.Fprintf(writer, "%v\n", 5)
	}
}
