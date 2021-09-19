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
	var d = make([]int, n)
	var s = make([]int, n)
	var t = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &t[i])
	}

	var p = make([]pair, n)
	for i := 0; i < n; i++ {
		p[i].first = d[i]*60 + s[i]
		p[i].second = d[i]*60 + t[i]
	}
	sort.Slice(p, func(i, j int) bool { return p[i].second < p[j].second })

	lst := -1 << 60
	for i := 0; i < n; i++ {
		if lst <= p[i].first {
			lst = p[i].second
		} else {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")
}
