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

	var a float64
	fmt.Fscan(reader, &a)
	var b float64
	fmt.Fscan(reader, &b)

	x := (a + b) / 2
	y := (a - b) / 2

	fmt.Fprintf(writer, "%v %v\n", x, y)
}
