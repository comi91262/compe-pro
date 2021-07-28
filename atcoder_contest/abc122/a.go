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
	var b string
	fmt.Fscan(reader, &b)

	switch b {
	case "A":
		fmt.Fprintf(writer, "%v\n", "T")
	case "C":
		fmt.Fprintf(writer, "%v\n", "G")
	case "G":
		fmt.Fprintf(writer, "%v\n", "C")
	case "T":
		fmt.Fprintf(writer, "%v\n", "A")
	}
}
