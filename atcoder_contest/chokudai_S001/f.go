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
	c := make([]int, n)
	copy(c, a)
	for i := 1; i < n; i++ {
		c[i] = max(c[i], c[i-1])
	}
	ans := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			ans++
			continue
		}
		if c[i-1] < a[i] {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
