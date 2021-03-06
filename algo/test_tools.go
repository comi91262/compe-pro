package algo

import (
	"errors"
)

func testEqSlice(a, b []int) bool {

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

func testEqSlice2(a, b [][]int) bool {

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
			if !testEqSlice(a[i], b[i]) {
				return false
			}
		}
	}

	return true
}

func testEqMatrix(a, b [][]int) (bool, error) {
	if len(a) != len(b) || len(a) == 0 {
		return false, errors.New("row size error")
	}

	n := len(a)
	for i := 0; i < n; i++ {
		if len(a[i]) != len(b[i]) {
			return false, errors.New("col size error")
		}
	}

	m := len(a[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] != b[i][j] {
				return false, nil
			}
		}
	}

	return true, nil
}
