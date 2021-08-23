package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, n+1)

	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		l--
		r--

		a[l]++
		a[r+1]--
	}

	for i := 1; i < n+1; i++ {
		a[i] = a[i] + a[i-1]
	}

	ans := 0
	for i := 0; i < n+1; i++ {
		if a[i] == m {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
