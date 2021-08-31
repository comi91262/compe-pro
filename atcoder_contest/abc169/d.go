package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func divisor(n int) []int {
	var r []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			r = append(r, i)
			if i*i != n {
				r = append(r, n/i)
			}
		}
	}
	return r
}

func primeFactor(n int) map[int]int {
	var m = map[int]int{}

	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n]++
	}
	return m
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	d := divisor(n)
	sort.Ints(d)
	// fmt.Fprintf(writer, "%v\n", d)

	ans := 0
	for i := 0; i < len(d); i++ {
		if d[i] == 1 {
			continue
		}
		m := primeFactor(d[i])
		if n%d[i] == 0 && len(m) == 1 {
			ans++
			n /= d[i]
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
