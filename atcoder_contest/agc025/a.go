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

	ans := 1 << 60
	for i := 1; i <= n-1; i++ {
		a, b := i, n-i

		ad := 0
		bd := 0
		for a > 0 {
			ad += a % 10
			a /= 10
		}
		for b > 0 {
			bd += b % 10
			b /= 10
		}
		ans = min(ans, bd+ad)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
