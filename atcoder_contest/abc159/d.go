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

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	total := 0
	for _, v := range m {
		total += combination(v, 2)
	}

	for i := 0; i < n; i++ {
		f := combination(m[a[i]], 2)
		l := combination(m[a[i]]-1, 2)
		fmt.Fprintf(writer, "%v\n", total-f+l)
	}

}
