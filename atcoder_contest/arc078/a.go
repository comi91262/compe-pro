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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			b[i] = a[i]
			continue
		}
		b[i] = b[i-1] + a[i]
	}

	var c = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			c[i] = a[n-1-i]
			continue
		}
		c[i] = c[i-1] + a[n-1-i]
	}

	ans := 1 << 60
	for i := 0; i < n-1; i++ {
		ans = min(ans, abs(b[i]-c[n-2-i]))
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
