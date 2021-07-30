package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func permutation(n int, k int) int {
	if n < k {
		return 0
	}

	prod := 1
	for k > 0 {
		prod *= n - k + 1
		k--
	}
	return prod
}

func combination(n int, k int) int {
	return permutation(n, k) / permutation(k, k)
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 0; i < n; i++ {
		a[i] %= 200
	}

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	ans := 0
	for _, v := range m {
		ans += combination(v, 2)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
