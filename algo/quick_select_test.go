package algo

import "testing"

func TestInsertionSort(t *testing.T) {
	a := []int{8, 7, 6, 5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}

	insertionSort(a)

	if !testEqSlice(a, want) {
		t.Errorf("Error 0: %v", a)
	}
}
