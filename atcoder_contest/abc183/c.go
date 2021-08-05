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

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// スライス a の i番目からj番目を反転させる関数
func reverse(a []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}
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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &t[i][j])
		}
	}

	a := []int{}
	for i := 1; i < n; i++ {
		a = append(a, i+1)
	}

	ans := 0
	for i := 0; i < permutation(n-1, n-1); i++ {
		// fmt.Fprintf(writer, "%v\n", a)
		pre := 1
		cost := 0
		for j := range a {
			cost += t[pre-1][a[j]-1]
			pre = a[j]
		}
		if cost+t[pre-1][0] == k {
			ans++
		}

		nextPermutation(a)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
