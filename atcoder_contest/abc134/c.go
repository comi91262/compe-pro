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

	var b = make([]int, n)
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i]
		c[i] = a[i]
	}
	for i := 1; i < n; i++ {
		b[i] = max(b[i], b[i-1])
	}
	for i := 0; i < n-1; i++ {
		c[n-i-2] = max(c[n-i-1], c[n-i-2])
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Fprintf(writer, "%v\n", c[1])
			continue
		}
		if i == n-1 {
			fmt.Fprintf(writer, "%v\n", b[n-2])
			continue
		}
		fmt.Fprintf(writer, "%v\n", max(c[i+1], b[i-1]))
	}
}
