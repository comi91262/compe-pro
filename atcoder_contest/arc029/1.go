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
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var t = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}

	ans := 1 << 60
	for i := 0; i < 1<<n; i++ {
		a := 0
		b := 0
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				a += t[j]
			} else {
				b += t[j]
			}
		}
		ans = min(ans, max(a, b))
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
