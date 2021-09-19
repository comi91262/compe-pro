package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func combination(n, k int) int {
	r := 1
	for d := 1; d <= k; d++ {
		r *= n
		n--
		r /= d
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
	}

	b := []int{0}
	b = append(b, a...)
	for i := 0; i < n; i++ {
		b[i+1] += b[i]
	}
	m := map[int]int{}
	for i := 0; i < n+1; i++ {
		m[b[i]]++
	}
	ans := 0
	for _, v := range m {
		if v >= 2 {
			ans += combination(v, 2)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
