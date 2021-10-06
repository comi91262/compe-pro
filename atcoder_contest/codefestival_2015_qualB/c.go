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
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })

	cur := 0

	for i := 0; i < m; i++ {
		if cur < n && a[cur] >= b[i] {
			cur++
		} else {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "YES")
}
