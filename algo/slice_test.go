package algo

import (
	"testing"
)

func TestReverse(t *testing.T) {
	got1 := []int{1, 2}
	reverse(got1, 0, len(got1)-1)
	if !testEqSlice(got1, []int{2, 1}) {
		t.Errorf("error: %v", got1)
	}

	got2 := []int{3, 1, 2}
	reverse(got2, 0, len(got2)-1)
	if !testEqSlice(got2, []int{2, 1, 3}) {
		t.Errorf("error: %v", got2)
	}
}

func TestMex(t *testing.T) {
	if got := mex([]int{}); got != 0 {
		t.Errorf("error: %v", got)
	}
	if got := mex([]int{0}); got != 1 {
		t.Errorf("error: %v", got)
	}
	if got := mex([]int{1}); got != 0 {
		t.Errorf("error: %v", got)
	}
	if got := mex([]int{3, 1, 2, 0}); got != 4 {
		t.Errorf("error: %v", got)
	}
	if got := mex([]int{0, 1, 0}); got != 2 {
		t.Errorf("error: %v", got)
	}
	if got := mex([]int{2, 0, 0}); got != 1 {
		t.Errorf("error: %v", got)
	}
}
