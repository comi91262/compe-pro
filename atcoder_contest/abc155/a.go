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

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	if a == b && a != c {
		fmt.Fprintf(writer, "%v\n", "Yes")
		return

	}

	if c == b && a != b {
		fmt.Fprintf(writer, "%v\n", "Yes")
		return

	}
	if c == a && c != b {
		fmt.Fprintf(writer, "%v\n", "Yes")
		return

	}

	fmt.Fprintf(writer, "%v\n", "No")
}
