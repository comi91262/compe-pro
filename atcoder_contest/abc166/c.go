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
	var h = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}

	var g = make([][]int, n)
	var a = make([]int, m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}
	ans := 0

	for i := 0; i < n; i++ {
		flag := true
		for _, j := range g[i] {
			if h[i] <= h[j] {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
