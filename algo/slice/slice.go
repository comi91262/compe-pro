package slice

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// スライス a の i番目からj番目を反転させる関数
func reverse(a []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}

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

// a, b がソート済
func intersect(a, b []int) []int {
	r := []int{}
	idx := 0
	for i := 0; i < len(a); i++ {
		if idx < len(b) && a[i] == b[idx] {
			r = append(r, a[i])
			idx++
		}
	}
	return r
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
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
func mex(arg ...int) int {
	m := map[int]int{}
	for _, x := range arg {
		m[x]++
	}

	for i := 0; ; i++ {
		if m[i] == 0 {
			return i
		}
	}
}
