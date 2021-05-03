package main

import (
	"fmt"
)

func uniq(a []int) []int {
	m := make(map[int]bool)
	result := []int{}

	for _, e := range a {
		if !m[e] {
			m[e] = true
			result = append(result, e)
		}
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	// sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	fmt.Println(len(uniq(a)))
}
