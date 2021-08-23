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
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
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

	var w = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &w[i])
	}

	s1 := 0
	s2 := sum(w)
	ans := 1 << 60
	for i := 0; i < n; i++ {
		s1 += w[i]
		s2 -= w[i]

		ans = min(ans, abs(s1-s2))
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
