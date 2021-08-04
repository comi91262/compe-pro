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

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var u string
	fmt.Fscan(reader, &u)

	if s == u {
		fmt.Fprintf(writer, "%v %v\n", a-1, b)
	} else {
		fmt.Fprintf(writer, "%v %v\n", a, b-1)
	}
}
