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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i]
	}
	for i := 0; i < n-1; i++ {
		b[i+1] += b[i]
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		ans += (n-i-1)*a[i] - (b[n-1] - b[i])
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
