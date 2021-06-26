package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func isOK(a *[]int, index, key int) bool {
	return (*a)[index] >= key
}

func binarySearch(a *[]int, key int) int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		} else {
			return x * -1
		}
	}

	ng := -1
	ok := len(*a)

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if isOK(a, mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n*2)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := n; i < 2*n; i++ {
		a[i] = a[i-n]
	}

	b := make([]int, n*2)
	for i := 0; i < 2*n; i++ {
		if i == 0 {
			b[i] = a[i]
			continue
		}
		b[i] = a[i] + b[i-1]
	}
	bn := b[n-1] / 10

	for i := 0; i < n; i++ {
		br := b[i] + bn
		idx := binarySearch(&b, br)

		if 10*(b[idx]-b[i]) == b[n-1] {
			fmt.Fprintf(writer, "Yes\n")
			return
		}
	}

	fmt.Fprintf(writer, "No\n")
}
