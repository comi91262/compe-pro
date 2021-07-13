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

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]int, n)
	var b = make([]int, 2*n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	for i := n; i < 2*n; i++ {
		b[i] = a[i-n] - b[i-n]
	}

	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	ans := 0
	for i := 0; i < k; i++ {
		ans += b[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
