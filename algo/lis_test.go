package algo

import "testing"

func TestLIS0(t *testing.T) {
	var a = []int{3, 1, 4, 1, 5, 9, 2, 6}
	got := lis(a)
	want := []int{1, 1, 2, 1, 3, 4, 2, 4}

	if !testEqSlice(got, want) {
		t.Errorf("Error 0: %v", got)
	}
}
