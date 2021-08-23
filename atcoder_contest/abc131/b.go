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
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var l int
	fmt.Fscan(reader, &l)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = l + i
	}

	total := sum(a)
	mn := 1 << 60
	for i := 0; i < n; i++ {
		mn = min(abs(a[i]), mn)
	}

	for i := 0; i < n; i++ {
		if mn == abs(a[i]) {
			fmt.Fprintf(writer, "%v\n", total-a[i])
			break

		}
	}

}
