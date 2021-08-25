package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func fact(n int) int {
	r := 1
	for n > 0 {
		r *= n
		n--
	}
	return r
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// スライス a の i番目からj番目を反転させる関数
func reverse(a []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}

// Narayana Pandita’s algorithm
// swap, reverse (slice.go) が必要
func nextPermutation(a []int) {
	var lastIndex = len(a) - 1

	if lastIndex < 1 {
		return
	}

	var i = lastIndex - 1
	for i >= 0 && a[i] >= a[i+1] {
		i -= 1
	}
	if i < 0 {
		reverse(a, 0, lastIndex)
		return
	}

	var j = lastIndex
	for j > i+1 && a[j] <= a[i] {
		j -= 1
	}

	swap(a, i, j)
	reverse(a, i+1, lastIndex)
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}

	}
	return true
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}
	var q = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &q[i])
	}

	c := []int{}
	for i := 0; i < n; i++ {
		c = append(c, i+1)
	}

	a := 0
	b := 0
	for i := 0; i < fact(n); i++ {
		if eq(p, c) {
			a = i
		}

		if eq(q, c) {
			b = i
		}

		nextPermutation(c)
	}

	fmt.Fprintf(writer, "%v\n", abs(a-b))
}
