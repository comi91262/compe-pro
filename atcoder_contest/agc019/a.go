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
	var q int
	fmt.Fscan(reader, &q)
	var h int
	fmt.Fscan(reader, &h)
	var s int
	fmt.Fscan(reader, &s)
	var d int
	fmt.Fscan(reader, &d)
	var n int
	fmt.Fscan(reader, &n)

	mn := min(4*q, 2*h, s)

	if mn*2 < d {
		fmt.Fprintf(writer, "%v\n", n*mn)
	} else {
		if n%2 == 0 {
			fmt.Fprintf(writer, "%v\n", n/2*d)
		} else {
			fmt.Fprintf(writer, "%v\n", n/2*d+mn)
		}

	}

}

