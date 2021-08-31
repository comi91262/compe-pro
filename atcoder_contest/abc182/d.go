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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	b := make([]int, len(a))
	copy(b, a)
	for i := 0; i < n-1; i++ {
		b[i+1] += b[i]
	}
	c := make([]int, len(a))
	copy(c, b)
	for i := 0; i < n-1; i++ {
		c[i+1] += c[i]
	}
	d := make([]int, len(a))
	for i := 0; i < n; i++ {
		if i == 0 {
			d[i] = b[i]
			continue
		}
		d[i] = max(d[i-1], b[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			ans = max(ans, d[i])
			continue
		}
		ans = max(ans, d[i]+c[i-1])
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
