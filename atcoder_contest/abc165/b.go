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

	a := 100
	y := 0
	for x > a {
		y++
		a += a / 100
	}
	fmt.Fprintf(writer, "%v\n", y)

}
