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
func calc(x int, a []int) int {
	r := 0
	for i := range a {
		r += abs(a[i] - (x + i + 1))
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

	ans, cl, cr := 1<<60, -(1 << 40), 1<<40

	for abs(cl-cr) > 2 {

		c1 := (cl + cl + cr) / 3
		c2 := (cl + cr + cr) / 3

		d1 := calc(c1, a)
		d2 := calc(c2, a)

		ans = min(ans, d1, d2)
		if d1 > d2 {
			cl = c1
		} else {
			cr = c2
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
