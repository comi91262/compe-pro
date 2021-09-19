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
	var q int
	fmt.Fscan(reader, &q)

	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = n
	}
	b := make([][]int, n)
	for i := 0; i < q; i++ {
		var t, n, m int
		fmt.Fscan(reader, &t)
		switch t {
		case 1:
			fmt.Fscan(reader, &n)
			n--
			sum := 0
			for j := range b[n] {
				sum += a[b[n][j]]
			}
			fmt.Fprintf(writer, "%v\n", sum)
		case 2:
			fmt.Fscan(reader, &n)
			fmt.Fscan(reader, &m)
			n--
			m--
			a[m]--
			b[n] = append(b[n], m)
		}
	}
}
