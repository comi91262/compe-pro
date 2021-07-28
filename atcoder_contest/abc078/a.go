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

	var x string
	fmt.Fscan(reader, &x)
	var y string
	fmt.Fscan(reader, &y)

	if x < y {
		fmt.Fprintf(writer, "%v\n", "<")
	} else if x == y {
		fmt.Fprintf(writer, "%v\n", "=")
	} else {
		fmt.Fprintf(writer, "%v\n", ">")
	}
}
