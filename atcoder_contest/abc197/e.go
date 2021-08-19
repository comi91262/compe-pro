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

type pair struct {
	x int
	y int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	mn := make([]int, n)
	mx := make([]int, n)
	for i := 0; i < n; i++ {
		mn[i] = 1 << 60
		mx[i] = -1 << 60
	}

	for i := 0; i < n; i++ {
		var x, c int
		fmt.Fscan(reader, &x, &c)
		c--
		mn[c] = min(mn[c], x)
		mx[c] = max(mx[c], x)
	}

	d := []pair{{0, 0}}
	for i := 0; i < n; i++ {
		if mn[i] == 1<<60 {
			continue
		}
		d = append(d, pair{mn[i], mx[i]})
	}
	d = append(d, pair{0, 0})

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(d))
		for j := 0; j < len(d); j++ {
			dp[i][j] = 1 << 60
		}
	}

	dp[0][0], dp[1][0] = 0, 0
	for i := 1; i < len(d); i++ {
		x, y := d[i].x, d[i].y
		px, py := d[i-1].x, d[i-1].y

		dp[0][i] = min(dp[0][i-1]+abs(px-y)+abs(y-x), dp[1][i-1]+abs(py-y)+abs(y-x))
		dp[1][i] = min(dp[0][i-1]+abs(px-x)+abs(x-y), dp[1][i-1]+abs(py-x)+abs(x-y))
	}

	fmt.Fprintf(writer, "%v\n", min(dp[0][len(d)-1], dp[1][len(d)-1]))
}
