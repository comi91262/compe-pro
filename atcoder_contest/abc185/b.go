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
	var t int
	fmt.Fscan(reader, &t)

	var a = make([]int, m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
	}
	maxN := n

	for i := 0; i < m; i++ {
		if i == 0 {
			n -= a[i]
		} else {
			n -= (a[i] - b[i-1])
		}

		if n <= 0 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
		n = min(n+b[i]-a[i], maxN)
	}
	n -= (t - b[m-1])
	if n <= 0 {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
}
