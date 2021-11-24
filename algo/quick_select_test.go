package algo

import "testing"

func TestInsertionSort(t *testing.T) {
	a := []int{8, 7, 6, 5, 4, 3, 2, 1}

	insertionSort(a)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !testEqSlice(a, want) {
		t.Errorf("Error 0: %v", a)
	}
}

func TestMedian0(t *testing.T) {
	a := []int{8, 7, 6, 5, 4, 3, 2, 1}
	got := median(a)
	want := 4

	if got != want {
		t.Errorf("Error 0: %v", got)
	}
}

func TestMedian1(t *testing.T) {
	a := []int{8, 7, 6, 5, 4, 3, 2}
	got := median(a)
	want := 5

	if got != want {
		t.Errorf("Error 0: %v", got)
	}
}
