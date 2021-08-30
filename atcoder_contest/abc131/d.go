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
	first  int
	second int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var p = make([]pair, n)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		p[i].first = a[i]
		p[i].second = b[i]
	}
	sort.Slice(p, func(i, j int) bool { return p[i].second < p[j].second })

	t := 0
	for i := 0; i < n; i++ {
		t += p[i].first
		if t > p[i].second {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
}
