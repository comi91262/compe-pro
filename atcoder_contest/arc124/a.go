package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 998244353

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	var c = make([]string, k)
	var t = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &c[i], &t[i])
	}
	a := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		a[i] = map[int]int{}
	}
	b := make([]bool, n)

	for i := 0; i < k; i++ {
		switch c[i] {
		case "L":
			b[t[i]-1] = true
			for j := 0; j < t[i]-1; j++ {
				a[j][t[i]]++
			}
		case "R":
			b[t[i]-1] = true
			for j := t[i]; j < n; j++ {
				a[j][t[i]]++
			}
		}
	}
	//fmt.Fprintf(writer, "%v\n", a)
	//fmt.Fprintf(writer, "%v\n", b)

	ans := 1
	for i := 0; i < n; i++ {
		if !b[i] {
			ans *= (k - len(a[i]))
			ans %= mod
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
