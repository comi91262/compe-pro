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

	var g = make([][]int, n+1)
	var a = make([]int, m+1)
	var b = make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "%v\n", len(g[i]))
	}

}
