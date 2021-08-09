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

	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)
	var z int
	fmt.Fscan(reader, &z)

	x, y = y, x
	x, z = z, x

	fmt.Fprintf(writer, "%v %v %v\n", x, y, z)
}
