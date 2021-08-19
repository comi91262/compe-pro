package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	x int
	y int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]pair, n)
	for i := 0; i < n; i++ {
		a[i].x = i
		fmt.Fscan(reader, &a[i].y)
	}

	sort.Slice(a, func(i, j int) bool { return a[i].y < a[j].y })

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if i < n/2 {
			ans[a[i].x] = a[n/2].y
		} else {
			ans[a[i].x] = a[n/2-1].y
		}
	}

	for i := range ans {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}
}
