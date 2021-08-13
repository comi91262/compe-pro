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
	a int
	b int
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, n)
	var b = make([]int, n)
	var s = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		s[i].a, s[i].b = a[i], b[i]
	}
	sort.Slice(s, func(i, j int) bool { return s[i].a < s[j].a })

	ans := 0
	mm := 0
	for i := range s {
		for j := 0; j < s[i].b; j++ {
			ans += s[i].a
			mm++
			if mm == m {
				fmt.Fprintf(writer, "%v\n", ans)
				return
			}
		}
	}

}
