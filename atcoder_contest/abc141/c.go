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
	var k int
	fmt.Fscan(reader, &k)
	var q int
	fmt.Fscan(reader, &q)
	var a = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &a[i])
		a[i]--
	}

	m := map[int]int{}
	for i := 0; i < q; i++ {
		m[a[i]]++
	}

	for i := 0; i < n; i++ {
		if k-q+m[i] > 0 {
			fmt.Fprintf(writer, "%v\n", "Yes")
		} else {
			fmt.Fprintf(writer, "%v\n", "No")
		}
	}

}
