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

	var l = make([]int, q)
	var r = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &l[i])
		fmt.Fscan(reader, &r[i])
		l[i]--
		r[i]--
	}

	a := make([]int, n+1)
	for i := 0; i < q; i++ {
		a[l[i]]++
		a[r[i]+1]--
	}

	for i := 0; i < n; i++ {
		a[i+1] += a[i]
	}

	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			fmt.Fprintf(writer, "%v", "0")
		} else {
			fmt.Fprintf(writer, "%v", "1")
		}
	}
	fmt.Fprintf(writer, "\n")

}
