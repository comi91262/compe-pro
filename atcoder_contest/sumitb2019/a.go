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
	var m1 int
	fmt.Fscan(reader, &m1)
	var d1 int
	fmt.Fscan(reader, &d1)
	var m2 int
	fmt.Fscan(reader, &m2)
	var d2 int
	fmt.Fscan(reader, &d2)

	if m1 != m2 {
		fmt.Fprintf(writer, "%v\n", 1)
	} else {
		fmt.Fprintf(writer, "%v\n", 0)
	}
}
