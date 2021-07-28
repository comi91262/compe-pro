package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	fmt.Fprintf(writer, "%v\n", min(n*a, b))
}
