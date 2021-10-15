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
	mpa := map[int]int{}
	mpb := map[int]int{}
	for i := 0; i < n; i++ {
		mpa[a[i]]++
	}
	for i := 0; i < m; i++ {
		mpb[b[i]]++
	}
	ans := []int{}
	for k := range mpa {
		if mpb[k] > 0 {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)

	for i := range ans {
		if i > 0 {
			fmt.Fprintf(writer, " ")
		}
		fmt.Fprintf(writer, "%v", ans[i])
	}
	fmt.Fprintf(writer, "\n")

}
