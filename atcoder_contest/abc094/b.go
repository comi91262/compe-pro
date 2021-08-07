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
	var m int
	fmt.Fscan(reader, &m)
	var x int
	fmt.Fscan(reader, &x)

	var a = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var b = make([]int, n+1)
	for i := 0; i < m; i++ {
		b[a[i]]++
	}

	s := 0
	for i := x; i <= n; i++ {
		s += b[i]
	}

	t := 0
	for i := 0; i <= x; i++ {
		t += b[i]
	}

	fmt.Fprintf(writer, "%v\n", min(s, t))
}
