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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	total := sum(a)

	ave1 := total / n
	ave2 := (total + n - 1) / n

	sum1 := 0
	sum2 := 0
	for i := 0; i < n; i++ {
		diff1 := (a[i] - ave1)
		diff2 := (a[i] - ave2)

		sum1 += diff1 * diff1
		sum2 += diff2 * diff2
	}
	fmt.Fprintf(writer, "%v\n", min(sum1, sum2))
}
