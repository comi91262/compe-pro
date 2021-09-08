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
	var m int
	fmt.Fscan(reader, &m)

	a := make([]int, m+1)
	total := 0
	for i := 0; i < n; i++ {
		var l, r, s int
		fmt.Fscan(reader, &l, &r, &s)
		l--
		r--
		total += s
		a[l] += s
		a[r+1] -= s
	}
	for i := 0; i < m; i++ {
		a[i+1] += a[i]
	}

	mn := 1 << 60
	for i := 0; i < m; i++ {
		mn = min(mn, a[i])
	}
	fmt.Fprintf(writer, "%v\n", total-mn)
}
