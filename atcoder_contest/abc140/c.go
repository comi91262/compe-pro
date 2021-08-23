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

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var b = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &b[i])
	}

	var a = make([]int, n)
	a[0] = b[0]
	a[n-1] = b[n-2]
	for i := 1; i < n-1; i++ {
		a[i] = min(b[i-1], b[i])
	}

	fmt.Fprintf(writer, "%v\n", sum(a))
}
