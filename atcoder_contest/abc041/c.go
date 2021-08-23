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
		fmt.Fscan(reader, &a[i].x)
		a[i].y = i + 1
	}

	sort.Slice(a, func(i, j int) bool { return a[i].x > a[j].x })
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", a[i].y)
	}
}
