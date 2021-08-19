package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	if m == 0 {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}

	var a = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
		a[i]--
	}

	sort.Ints(a)
	var b = make([]int, m+1)
	for i := 0; i < m; i++ {
		if i == 0 {
			b[0] = a[i] - 0
			continue
		}
		b[i] = a[i] - a[i-1] - 1
	}
	b[m] = n - 1 - a[m-1]

	k := 1 << 60
	for i := range b {
		if b[i] == 0 {
			continue
		}
		k = min(k, b[i])
	}
	// fmt.Fprintf(writer, "%v\n", k)
	// fmt.Fprintf(writer, "%v\n", b)
	// fmt.Fprintf(writer, "%v\n", a)

	ans := 0
	for i := range b {
		ans += (b[i] + k - 1) / k
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
