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
	var t int
	fmt.Fscan(reader, &t)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	size := 2_000_000
	var b = make([]int, size)
	for i := 0; i < n; i++ {
		b[a[i]+1]++
		b[a[i]+t+1]--
	}
	// fmt.Fprintf(writer, "%v\n", b)

	for i := 0; i < size-1; i++ {
		b[i+1] += b[i]
	}

	ans := 0
	for i := 0; i < size; i++ {
		if b[i] > 0 {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
