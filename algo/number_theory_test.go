package algo

import (
	"testing"
)

func TestLcm0(t *testing.T) {
	got := lcm(4, 5)
	want := 20
	if got != want {
		t.Errorf("error: 0")
	}
}

func TestLcm1(t *testing.T) {
	got := lcm(4, 2)
	want := 4
	if got != want {
		t.Errorf("error: 1")
	}
}

func TestPermutation0(t *testing.T) {
	got := permutation(2, 3)
	want := 0
	if got != want {
		t.Errorf("error: 0")
	}
}

func TestPermutation1(t *testing.T) {
	got := permutation(3, 2)
	want := 6
	if got != want {
		t.Errorf("error: 1")
	}
}

func TestPermutation2(t *testing.T) {
	got := permutation(3, 0)
	want := 1
	if got != want {
		t.Errorf("error: 2")
	}
}

func TestCombination0(t *testing.T) {
	got := combination(2, 3)
	want := 0
	if got != want {
		t.Errorf("error: 0")
	}
}

func TestCombination1(t *testing.T) {
	got := combination(3, 2)
	want := 3
	if got != want {
		t.Errorf("error: 1")
	}
}

func TestCombination2(t *testing.T) {
	got := combination(3, 0)
	want := 1
	if got != want {
		t.Errorf("error: 2")
	}
}

func TestDivisor0(t *testing.T) {
	if !testSliceEq(divisor(36), []int{1, 36, 2, 18, 3, 12, 4, 9, 6}) {
		t.Errorf("error: 0")
	}
}
