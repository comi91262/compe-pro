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
	var k int
	fmt.Fscan(reader, &k)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	var b = []int{}
	for _, v := range m {
		b = append(b, v)
	}
	sort.Ints(b)

	ans := 0
	for i := 0; i < len(m)-k; i++ {
		ans += b[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
