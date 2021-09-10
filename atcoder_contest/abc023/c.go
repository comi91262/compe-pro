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
	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)
	var k int
	fmt.Fscan(reader, &k)
	var n int
	fmt.Fscan(reader, &n)

	var r = make([]int, n)
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &r[i])
		fmt.Fscan(reader, &c[i])
		r[i]--
		c[i]--
	}

	x := make([]int, h)
	y := make([]int, w)
	for i := 0; i < n; i++ {
		x[r[i]]++
		y[c[i]]++
	}
	xm := map[int]int{}
	for i := range x {
		xm[x[i]]++
	}
	ym := map[int]int{}
	for i := range y {
		ym[y[i]]++
	}

	ans := 0
	for i := 0; i <= k; i++ {
		ans += xm[i] * ym[k-i]
	}

	for i := 0; i < n; i++ {
		if x[r[i]]+y[c[i]] == k {
			ans--
		}
		if x[r[i]]+y[c[i]] == k+1 {
			ans++
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}

