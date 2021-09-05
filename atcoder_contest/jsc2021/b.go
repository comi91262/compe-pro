package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	ma := map[int]int{}
	for i := 0; i < n; i++ {
		ma[a[i]]++
	}
	for i := 0; i < m; i++ {
		ma[b[i]]++
	}
	ans := []int{}
	for k, cnt := range ma {
		if cnt > 1 {
			continue
		}
		ans = append(ans, k)
	}
	sort.Ints(ans)
	for i := 0; i < len(ans); i++ {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}

}
