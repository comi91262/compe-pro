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

	var a = make([][]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = make([]int, 10)
	}

	for i := 1; i <= n; i++ {
		top := -1
		btm := i % 10
		m := i
		for m > 0 {
			top = m % 10
			m /= 10
		}
		a[top][btm]++
	}
	ans := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ans += a[i][j] * a[j][i]
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
