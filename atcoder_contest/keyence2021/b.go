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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}
	m[0] = min(m[0], k)

	ans := 0
	for i := 1; i < n; i++ {
		if m[i] > m[i-1] {
			m[i] = m[i-1]
		} else {
			ans += min(k, m[i-1]-m[i]) * i
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
