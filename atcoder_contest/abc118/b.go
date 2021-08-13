package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		var k int
		fmt.Fscan(reader, &k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	ma := map[int]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ma[a[i][j]]++
		}
	}
	// fmt.Fprintf(writer, "%v\n", ma)

	ans := 0
	for k, v := range ma {
		if k == 0 {
			continue
		}
		if v == n {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
