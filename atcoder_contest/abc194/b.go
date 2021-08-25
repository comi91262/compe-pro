package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	x int
	y int
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
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
	}

	ans := 1 << 60
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				ans = min(ans, a[i]+b[j])
				continue
			}

			ans = min(ans, max(a[i], b[j]))
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
