package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i]--
	}
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	var c = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &c[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += b[i]

		if i != 0 && a[i-1]+1 == a[i] {
			ans += c[a[i-1]]
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
