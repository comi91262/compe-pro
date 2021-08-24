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

	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}
	sort.Ints(p)

	ans := 0
	for i := 0; i < k; i++ {
		ans += p[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
