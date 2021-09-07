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
	var m int
	fmt.Fscan(reader, &m)

	var p = make([]pair, m)
	var a = make([]int, m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		p[i].first = a[i]
		p[i].second = b[i]
	}
	sort.Slice(p, func(i, j int) bool { return p[i].second < p[j].second })

	ans := 0
	end := 0
	for i := 0; i < m; i++ {
		if end <= p[i].first {
			end = p[i].second
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
