package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

	var v = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &v[i])
	}
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	ans := 0
	for i := 0; i < 1<<n; i++ {
		x, y := 0, 0
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				x += v[j]
				y += c[j]

			}
		}

		ans = max(ans, x-y)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}

