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
	var w int
	fmt.Fscan(reader, &w)

	var s = make([]int, n)
	var t = make([]int, n)
	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &t[i])
		fmt.Fscan(reader, &p[i])
	}

	size := 200001
	var a = make([]int, size)
	for i := 0; i < n; i++ {
		a[s[i]] += p[i]
		a[t[i]] -= p[i]
	}
	for i := 0; i < size-1; i++ {
		a[i+1] += a[i]
	}

	for i := 0; i < size; i++ {
		if a[i] > w {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
