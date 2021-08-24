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
	var l = make([]int, n)
	var r = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i], &r[i])
	}

	size := 1_000_002
	c := make([]int, size)
	for i := 0; i < n; i++ {
		c[l[i]]++
		c[r[i]+1]--
	}
	for i := 1; i < size; i++ {
		c[i] += c[i-1]
	}

	ans := 0
	for i := 0; i < size; i++ {
		ans = max(ans, c[i])
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
