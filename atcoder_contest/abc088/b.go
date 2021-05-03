package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	aliceTotal := 0
	bobTotal := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			aliceTotal += a[i]
		} else {
			bobTotal += a[i]
		}
	}

	fmt.Println(aliceTotal - bobTotal)
}
