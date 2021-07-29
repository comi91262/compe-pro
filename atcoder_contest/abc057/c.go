package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func divisor(n int) []int {
	var r []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			r = append(r, i)
			if i*i != n {
				// r = append(r, n/i)
			}
		}
	}
	return r
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func f(a, b int) int {
	an := 0
	for a > 0 {
		an++
		a /= 10
	}

	bn := 0
	for b > 0 {
		bn++
		b /= 10
	}

	return max(an, bn)
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	d := divisor(n)
	sort.Ints(d)

	ans := 1 << 60
	for i := range d {
		ans = min(ans, f(d[i], n/d[i]))
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
