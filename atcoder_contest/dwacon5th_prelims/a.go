package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	first  int
	second int
}

func sum(a []pair) int {
	r := 0
	for i := range a {
		r += a[i].first
	}
	return r
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var a = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i].first)
		a[i].first *= n
		a[i].second = i
	}

	ans := 0
	mid := sum(a) / n
	mn := 1 << 60
	for i := 0; i < n; i++ {
		if mn > abs(mid-a[i].first) {
			ans = a[i].second
			mn = abs(mid - a[i].first)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
