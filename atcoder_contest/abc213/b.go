package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type data struct {
	idx   int
	point int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]data, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i].point)
		a[i].idx = i + 1
	}

	sort.Slice(a, func(i, j int) bool { return a[i].point > a[j].point })
	fmt.Fprintf(writer, "%v\n", a[1].idx)
}
