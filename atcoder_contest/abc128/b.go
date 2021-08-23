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
	i int
	s string
	p int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var p = make([]pair, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i].s)
		fmt.Fscan(reader, &p[i].p)
		p[i].i = i + 1
	}

	sort.Slice(p, func(i, j int) bool {
		if p[i].s == p[j].s {
			return p[i].p > p[j].p
		}
		return p[i].s < p[j].s
	})

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", p[i].i)
	}
}
