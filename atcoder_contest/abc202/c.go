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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i]--
	}
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
		b[i]--
	}
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
		c[i]--
	}
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[b[c[i]]]++
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += m[a[i]]
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
