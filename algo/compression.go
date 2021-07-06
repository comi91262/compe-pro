package algo

import "sort"

// compression 一次元座標圧縮
//
// unique(slice.go)とbinarySearch(search.go)が必要
func compression(a []int) []int {
	var n = len(a)
	var p = make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = a[i]
	}

	sort.Ints(p)
	p = unique(p)

	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = binarySearch(p, a[i], lowerBound)
	}

	return r
}
