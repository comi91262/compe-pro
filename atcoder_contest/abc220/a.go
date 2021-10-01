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

	d := (a + c - 1) / c * c
	if d <= b {
		fmt.Fprintf(writer, "%v\n", d)
	} else {
		fmt.Fprintf(writer, "%v\n", -1)
	}
}
