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
	var q int
	fmt.Fscan(reader, &q)

	var a = make([]int, n)

	for i := 0; i < q; i++ {
		var l int
		fmt.Fscan(reader, &l)
		var r int
		fmt.Fscan(reader, &r)
		var t int
		fmt.Fscan(reader, &t)

		for j := l; j <= r; j++ {
			a[j-1] = t
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", a[i])
	}
}
