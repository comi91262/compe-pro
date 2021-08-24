package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	var h = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}

	sort.Slice(h, func(i, j int) bool { return h[i] > h[j] })

	ans := 0
	for i := k; i < n; i++ {
		ans += h[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
