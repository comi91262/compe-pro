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

	var a1 int
	fmt.Fscan(reader, &a1)
	var a2 int
	fmt.Fscan(reader, &a2)
	var a3 int
	fmt.Fscan(reader, &a3)

	a := []int{a1, a2, a3}
	for i := 0; i < 6; i++ {
		if a[2]-a[1] == a[1]-a[0] {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}

		nextPermutation(a)
	}
	fmt.Fprintf(writer, "%v\n", "No")
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

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// スライス a の i番目からj番目を反転させる関数
func reverse(a []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}
