package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func isOK(a *[]int, index, key int) bool {
	return (*a)[index] > key
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

	var n, k, p int
	fmt.Fscan(reader, &n, &k, &p)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	mid := n / 2
	var b1 = make([][]int, n+1)
	var b2 = make([][]int, n+1)

	for i := 0; i < 1<<mid; i++ {
		sum := 0
		cnt := 0
		for j := 0; j < mid; j++ {
			if i&(1<<j) != 0 {
				sum += a[j]
				cnt++
			}
		}
		b1[cnt] = append(b1[cnt], sum)
	}

	for i := 0; i < 1<<(n-mid); i++ {
		sum := 0
		cnt := 0
		for j := 0; j < n-mid; j++ {
			if i&(1<<j) != 0 {
				sum += a[j+mid]
				cnt++
			}
		}
		b2[cnt] = append(b2[cnt], sum)
	}

	for i := 0; i <= n; i++ {
		sort.Ints(b1[i])
		sort.Ints(b2[i])
	}

	sum := 0
	for i := 0; i < k+1; i++ {
		for _, v := range b1[i] {
			idx := binarySearch(&b2[k-i], p-v)
			sum += idx
		}
	}

	fmt.Fprintf(writer, "%v\n", sum)
}
