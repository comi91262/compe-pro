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
	var x int
	fmt.Fscan(reader, &x)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	y := x
	m := 0
	for y > 0 {
		m++
		y /= 2
	}

	ans := 0
	for j := 0; j < m; j++ {
		if x&(1<<j) != 0 {
			ans += a[j]
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
