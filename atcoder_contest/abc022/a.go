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
	var s int
	fmt.Fscan(reader, &s)
	var t int
	fmt.Fscan(reader, &t)
	var w int
	fmt.Fscan(reader, &w)

	var a = make([]int, n+1)
	for i := 2; i < n+1; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	if s <= w && w <= t {
		ans++
	}

	for i := 2; i < n+1; i++ {
		w += a[i]
		if s <= w && w <= t {
			ans++
		}

	}

	fmt.Fprintf(writer, "%v\n", ans)
}
