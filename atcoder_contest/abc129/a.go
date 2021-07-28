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
	var p int
	fmt.Fscan(reader, &p)
	var q int
	fmt.Fscan(reader, &q)
	var r int
	fmt.Fscan(reader, &r)

	fmt.Fprintf(writer, "%v\n", min(p+q, q+r, r+p))
}
