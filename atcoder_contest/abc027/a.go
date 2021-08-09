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

	var l1 int
	fmt.Fscan(reader, &l1)
	var l2 int
	fmt.Fscan(reader, &l2)
	var l3 int
	fmt.Fscan(reader, &l3)

	if l1 == l2 {
		fmt.Fprintf(writer, "%v\n", l3)
		return
	}
	if l1 == l3 {
		fmt.Fprintf(writer, "%v\n", l2)
		return
	}
	if l2 == l3 {
		fmt.Fprintf(writer, "%v\n", l1)
		return
	}
}
