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

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 1 << 60
	for i := 0; i < 1<<(n-1); i++ {
		xored := 0
		ored := 0
		for j := 0; j <= n; j++ {
			if j < n {
				ored |= a[j]
			}

			if j == n || i>>j&1 == 1 {
				xored ^= ored
				ored = 0
			}
		}
		ans = min(ans, xored)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
