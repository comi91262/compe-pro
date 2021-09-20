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
	first  string
	second int
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)
	m := map[string]int{}
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)
		m[s]++
	}
	a := []pair{}
	for k, v := range m {
		a = append(a, pair{k, v})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].second > a[j].second })

	cnt := 0
	for i := 0; i < len(a); i++ {
		if a[i].second == a[k-1].second {
			cnt++
		}
	}
	if cnt > 1 {
		fmt.Fprintf(writer, "%v\n", "AMBIGUOUS")
		return
	}
	fmt.Fprintf(writer, "%v\n", a[k-1].first)
}
