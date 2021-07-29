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
	var t = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &t[i])
	}

	var m int
	fmt.Fscan(reader, &m)
	var p = make([]int, m+1)
	var x = make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(reader, &p[i])
		fmt.Fscan(reader, &x[i])
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += t[i]
	}

	for i := 1; i <= m; i++ {
		fmt.Fprintf(writer, "%v\n", sum-t[p[i]]+x[i])
	}

}

