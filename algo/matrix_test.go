package algo

import (
	"testing"
)

func TestProdMatrix(t *testing.T) {
	var a = [][]int{
		{1, 3},
		{2, 4},
	}
	var b = [][]int{
		{4, 1},
		{3, 2},
	}

	got := prodMatrix(a, b)

	want := [][]int{
		{13, 7},
		{20, 10},
	}

	if b, err := testEqMatrix(got, want); err != nil {
		t.Errorf("Error 0: %v %v %v", err, got, want)
	} else if !b {
		t.Errorf("Error 1: %v %v", got, want)
	}
}

func TestPowMatrix(t *testing.T) {
	var a = [][]int{
		{1, 1},
		{1, 0},
	}

	got := powMatrix(a, 8)
	want := [][]int{
		{34, 21},
		{21, 13},
	}

	if b, err := testEqMatrix(got, want); err != nil {
		t.Errorf("Error 0: %v", err)
	} else if !b {
		t.Errorf("Error 1: %v %v", got, want)
	}
}

func TestBitMatrix(t *testing.T) {
	var a = [][]int{
		{1, 0, 1, 1, 0},
		{1, 0, 0, 1, 1},
		{0, 1, 1, 1, 0},
		{0, 0, 1, 0, 1},
	}

	var got = sweepBitMatrix(a)
	var want = [][]int{
		{1, 0, 0, 1, 1},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0},
	}

	if b, err := testEqMatrix(got, want); err != nil {
		t.Errorf("Error 0: %v", err)
	} else if !b {
		t.Errorf("Error 1: %v %v", got, want)
	}

}
