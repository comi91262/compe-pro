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

func enumarate(size, init, inc int) []int {
	a := make([]int, size)
	if size == 0 {
		return a
	}

	a[0] = init
	for i := 1; i < size; i++ {
		a[i] = a[i-1] + inc
	}

	return a
}

// 最小除外数 (min-exclude) を求める
// 重複集合にも対応している
func mex(a []int) int {
	sort.Ints(a)
	mx := 0
	for _, x := range a {
		if x < mx {
			continue
		}
		if x != mx {
			break
		}
		mx++
	}
	return mx
}
