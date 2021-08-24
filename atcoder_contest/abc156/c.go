package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}
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
	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}
	ave := sum(x) / n

	ans1 := 0
	for i := 0; i < n; i++ {
		ans1 += (ave - x[i]) * (ave - x[i])
	}

	ans2 := 0
	for i := 0; i < n; i++ {
		ans2 += (ave + 1 - x[i]) * (ave + 1 - x[i])
	}

	fmt.Fprintf(writer, "%v\n", min(ans1, ans2))
}
