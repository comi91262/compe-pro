package algo

import "sort"

func unique(a []int) []int {
	r := make([]int, 0)
	m := make(map[int]bool, 0)

	for _, e := range a {
		if !m[e] {
			m[e] = true
			r = append(r, e)
		}
	}

	return r
}

// compression 一次元座標圧縮
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
		r[i] = sort.Search(len(p), func(j int) bool { return a[i] <= p[j] })
	}

	return r
}
