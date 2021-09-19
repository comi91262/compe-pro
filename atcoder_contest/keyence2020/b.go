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

	var x = make([]int, n)
	var l = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &l[i])
	}
	var a = make([]int, n)
	var b = make([]int, n)
	var p = make([]pair, n)
	for i := 0; i < n; i++ {
		a[i] = x[i] - l[i]
		b[i] = x[i] + l[i]
		p[i].first = a[i]
		p[i].second = b[i]
	}
	sort.Slice(p, func(i, j int) bool { return p[i].second < p[j].second })
	// fmt.Fprintf(writer, "%v\n", p)
	last := -1 << 60

	ans := 0
	for i := 0; i < n; i++ {
		if last <= p[i].first {
			last = p[i].second
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
