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
	case a == "H" && b == "H":
		fmt.Fprintf(writer, "%v\n", "H")
	case a == "H" && b == "D":
		fmt.Fprintf(writer, "%v\n", "D")
	case a == "D" && b == "H":
		fmt.Fprintf(writer, "%v\n", "D")
	case a == "D" && b == "D":
		fmt.Fprintf(writer, "%v\n", "H")
	}
}
