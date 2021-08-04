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

func calc(n, lim int) int {
	r := n - 1 - 2*max(n-1-lim, 0)
	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)
	k = abs(k)

	ans := 0
	for i := 2; i <= 2*n; i++ {
		ab := i
		cd := ab - k
		if ab < 2 || cd < 2 {
			continue
		}
		ans += calc(ab, n) * calc(cd, n)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
