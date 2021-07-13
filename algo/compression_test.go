package algo

import "testing"

func TestCompression0(t *testing.T) {
	var a = []int{1, 10, 5, 32, 99, 8, 10}
	got := compression(a)
	want := []int{0, 3, 1, 4, 5, 2, 3}

	if !testEqSlice(got, want) {
		t.Errorf("Error 0: %v", got)
	}
}
