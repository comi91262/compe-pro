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
	var k int
	fmt.Fscan(reader, &k)
	var p = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i].first)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i].second)
	}

	sort.Slice(p, func(i, j int) bool { return p[i].second < p[j].second })

	ans := 0
	c := map[int]int{}
	for i := 0; i < n; i++ {
		if c[p[i].first] > 0 {
			continue
		}
		ans += p[i].second
		c[p[i].first]++

		if len(c) == k {
			fmt.Fprintf(writer, "%v\n", ans)
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", -1)
}
