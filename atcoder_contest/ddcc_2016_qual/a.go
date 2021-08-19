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
	var c float64
	fmt.Fscan(reader, &c)
	fmt.Fprintf(writer, "%v\n", c*b/a)
}
