package algo

func testSliceEq(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func testSlice2Eq(a, b [][]int) bool {

	n := len(a)
	m := len(b)
	if n != m {
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i != j {
				continue
			}
			if !testSliceEq(a[i], b[i]) {
				return false
			}
		}
	}

	return true
}
