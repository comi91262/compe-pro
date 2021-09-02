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

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

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
	var c = make([]int, m)
	var p = make([]pair, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
		fmt.Fscan(reader, &c[i])
		p[i].first = b[i]
		p[i].second = c[i]
	}
	sort.Ints(a)
	sort.Slice(p, func(i, j int) bool { return p[i].second > p[j].second })

	idx := 0
	for i := 0; i < m; i++ {
		for idx < n && a[idx] < p[i].second && p[i].first > 0 {
			a[idx] = p[i].second
			p[i].first--
			idx++
		}

		if idx >= n {
			break
		}
	}
	fmt.Fprintf(writer, "%v\n", sum(a))

}
