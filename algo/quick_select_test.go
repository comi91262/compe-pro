package algo

import (
	"sort"
	"testing"
)

func TestInsertionSort1(t *testing.T) {
	a := []int{1, 3, 13, 76, 99, 8, 16, 20}
	insertionSort(a, 0, len(a)-1)
	want := []int{1, 3, 8, 13, 16, 20, 76, 99}
	if !testEqSlice(a, want) {
		t.Errorf("Error: %v", a)
	}
}

func TestInsertionSort2(t *testing.T) {
	a := []int{1, 13, 3, 76, 99, 8, 16, 20}
	insertionSort(a, 0, 2)
	want := []int{1, 3, 13, 76, 99, 8, 16, 20}
	if !testEqSlice(a, want) {
		t.Errorf("Error: %v", a)
	}
}

func TestMedian0(t *testing.T) {
	a := []int{8, 7, 6, 5, 4, 3, 2, 1}
	got := median(a, 0, len(a)-1)
	want := 4

	if got != want {
		t.Errorf("Error: %v", got)
	}
}

//func TestMedian1(t *testing.T) {
//	a := []int{8, 7, 6, 5, 4, 3, 2}
//	got := median(a, 1, 2)
//	want := 6
//
//	if got != want {
//		t.Errorf("Error: %v", got)
//	}
//}

//func TestSelectPivot0(t *testing.T) {
//	a := []int{8, 7, 6, 5, 4, 3, 2, 1}
//	got := selectPivot(a, 0, len(a)-1)
//	want := 6
//
//	if got != want {
//		t.Errorf("Error: %v", got)
//	}
//}

//func TestSelectPivot1(t *testing.T) {
//	a := []int{8, 7, 6, 5, 4, 3, 2, 1}
//	got := selectPivot(a, 1, 3)
//	want := 6
//
//	if got != want {
//		t.Errorf("Error: %v", got)
//	}
//}

func TestSelectPivot2(t *testing.T) {
	a := []int{1, 3, 13, 76, 99, 8, 16, 20, 25, 37, 7, 22, 28, 30, 47, 40, 43, 49, 52, 57, 5, 37, 70, 75, 85, 23, 33, 50, 77, 78}
	got := selectPivot(a, 0, len(a)-1)
	want := 28

	if got != want {
		t.Errorf("Error: %v", got)
	}
}

func TestQuickSelect(t *testing.T) {
	a := []int{1, 3, 13, 76, 99, 8, 16, 20, 25, 37, 7, 22, 28, 30, 47, 40, 43, 49, 52, 57, 5, 37, 70, 75, 85, 23, 33, 50, 77}
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(b)

	gots := []int{}
	for i := 0; i < len(a); i++ {
		gots = append(gots, QuickSelect(a, 0, len(a)-1, i))
	}

	if !testEqSlice(gots, b) {
		t.Errorf("Error 0: %v", b)
		t.Errorf("Error 0: %v", gots)
	}
}

func TestQuickSelect1(t *testing.T) {
	a := []int{1, 3, 13, 76, 99, 8, 16, 20, 25, 37, 7, 22, 28, 30, 47, 40, 43, 49, 52, 57, 5, 37, 70, 75, 85, 23, 33, 50, 77}
	b := make([]int, len(a))
	copy(b, a)
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })

	gots := []int{}
	for i := 0; i < len(a); i++ {
		gots = append(gots, QuickSelect(a, 0, len(a)-1, len(a)-i-1))
	}

	if !testEqSlice(gots, b) {
		t.Errorf("Error 0: %v", b)
		t.Errorf("Error 0: %v", gots)
	}
}
