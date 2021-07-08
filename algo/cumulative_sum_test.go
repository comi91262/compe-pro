package algo

import "testing"

func TestCumulativeSum0(t *testing.T) {

	var a = [][]int{
		[]int{2, 4, 6},
		[]int{4, 6, 2},
		[]int{1, 9, 1},
	}

	got := cumulativeSum2(a)
	want := [][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 2, 6, 12},
		[]int{0, 6, 16, 24},
		[]int{0, 7, 26, 35},
	}

	if !testSlice2Eq(got, want) {
		t.Errorf("Error 0: %v", got)
	}
}
