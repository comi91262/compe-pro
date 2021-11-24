package algo

import "testing"

func TestInsertionSort(t *testing.T) {
	a := []int{8, 7, 6, 5, 4, 3, 2, 1}

	insertionSort(a, 0, 1)
	want := []int{7, 8, 6, 5, 4, 3, 2, 1}
	if !testEqSlice(a, want) {
		t.Errorf("Error 0: %v", a)
	}

	insertionSort(a, 4, 5)
	want = []int{7, 8, 6, 5, 3, 4, 2, 1}
	if !testEqSlice(a, want) {
		t.Errorf("Error 1: %v", a)
	}

	insertionSort(a, 0, len(a)-1)
	want = []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !testEqSlice(a, want) {
		t.Errorf("Error 2: %v", a)
	}
}
