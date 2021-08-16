package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var s = make([]string, n)
	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &p[i])
	}

	total := sum(p)
	for i := 0; i < n; i++ {
		if 2*p[i] > total {
			fmt.Fprintf(writer, "%v\n", s[i])
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "atcoder")
}
