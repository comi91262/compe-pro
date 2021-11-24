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

func TestSelectPivot0(t *testing.T) {
	a := []int{8, 7, 6, 5, 4, 3, 2, 1}
	got := selectPivot(a, 0, len(a)-1)
	want := 6

	if got != want {
		t.Errorf("Error 0: %v", got)
	}
}

func TestSelectPivot1(t *testing.T) {
	a := []int{1, 3, 13, 76, 99, 8, 16, 20, 25, 37, 7, 22, 28, 30, 47, 40, 43, 49, 52, 57, 5, 37, 70, 75, 85, 23, 33, 50, 77, 78}
	got := selectPivot(a, 0, len(a)-1)
	want := 28

	if got != want {
		t.Errorf("Error 0: %v", got)
	}
}
