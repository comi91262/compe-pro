package algo

import (
	"testing"
)

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
