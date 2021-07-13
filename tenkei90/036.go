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

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()

	var n, Q int
	fmt.Fscan(reader, &n, &Q)

	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	var q = make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(reader, &q[i])
	}
	// rotate 45
	for i := 0; i < n; i++ {
		x[i], y[i] = x[i]-y[i], x[i]+y[i]
	}

	var xMin = 1_000_000_001
	var xMax = -1_000_000_001
	var yMin = 1_000_000_001
	var yMax = -1_000_000_001
	for i := 0; i < n; i++ {
		if xMin >= x[i] {
			xMin = x[i]
		}
		if yMin >= y[i] {
			yMin = y[i]
		}
		if xMax <= x[i] {
			xMax = x[i]
		}
		if yMax <= y[i] {
			yMax = y[i]
		}
	}

	for i := 0; i < Q; i++ {
		X, Y := x[q[i]-1], y[q[i]-1]
		fmt.Fprintf(writer, "%d\n", max(abs(xMin-X), abs(xMax-X), abs(yMin-Y), abs(yMax-Y)))
	}

}
