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
		var k int
		fmt.Fscan(reader, &k)
		a[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	var p int
	fmt.Fscan(reader, &p)
	var q int
	fmt.Fscan(reader, &q)
	var b = make([]int, p)
	for i := 0; i < p; i++ {
		fmt.Fscan(reader, &b[i])
	}
	ma := map[int]int{}
	for i := 0; i < p; i++ {
		ma[b[i]]++
	}
	ans := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := range a[i] {
			if ma[a[i][j]] > 0 {
				cnt++
			}
		}
		if cnt >= q {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
