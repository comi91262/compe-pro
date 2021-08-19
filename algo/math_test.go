package algo

import (
	"math"
	"testing"
)

func TestAbs0(t *testing.T) {
	got := abs(0)
	want := 0
	if got != want {
		t.Errorf("error: 0")
	}
}

// 32bit環境ではエラー
func TestAbs1(t *testing.T) {
	got := abs(math.MaxInt64)
	want := math.MaxInt64
	if got != want {
		t.Errorf("error: 1")
	}
}

func TestAbs2(t *testing.T) {
	got := abs(-100)
	want := 100
	if got != want {
		t.Errorf("error: 2")
	}
}

func TestMin1(t *testing.T) {
	got := min(0)
	want := 0
	if got != want {
		t.Errorf("error: 1")
	}
}

func TestMin2(t *testing.T) {
	got := min(0, 0)
	want := 0
	if got != want {
		t.Errorf("error: 2")
	}
}

func TestMin3(t *testing.T) {
	got := min(0, 1)
	want := 0
	if got != want {
		t.Errorf("error: 3")
	}
}

func TestMin4(t *testing.T) {
	got := min(1, 0)
	want := 0
	if got != want {
		t.Errorf("error: 4")
	}
}

func TestMin5(t *testing.T) {
	got := min(3, 2, 1)
	want := 1
	if got != want {
		t.Errorf("error: 5")
	}
}

func TestPow0(t *testing.T) {
	got := pow(3, 2)
	want := 9
	if got != want {
		t.Errorf("error: 0")
	}
}

func TestPow1(t *testing.T) {
	got := pow(3, 0)
	want := 1
	if got != want {
		t.Errorf("error: 1")
	}
}

func TestFloorLog0(t *testing.T) {
	got := floorLog(1)
	want := 0
	if got != want {
		t.Errorf("error: 0")
	}
}

func TestFloorLog1(t *testing.T) {
	got := floorLog(2)
	want := 1
	if got != want {
		t.Errorf("error: 1")
	}
}

func TestFloorLog2(t *testing.T) {
	got := floorLog(3)
	want := 2
	if got != want {
		t.Errorf("error: 2")
	}
}

func TestFloorLog3(t *testing.T) {
	got := floorLog(4)
	want := 2
	if got != want {
		t.Errorf("error: 3")
	}
}

func TestCeilLog0(t *testing.T) {
	got := ceilLog(1)
	want := 0
	if got != want {
		t.Errorf("error: 0")
	}
}

func TestCeilLog1(t *testing.T) {
	got := ceilLog(2)
	want := 1
	if got != want {
		t.Errorf("error: 1")
	}
}

func TestCeilLog2(t *testing.T) {
	got := ceilLog(3)
	want := 1
	if got != want {
		t.Errorf("error: 2")
	}
}

func TestCeilLog3(t *testing.T) {
	got := ceilLog(4)
	want := 2
	if got != want {
		t.Errorf("error: 3 %v", got)
	}
}

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
	if !testEqSlice(divisor(36), []int{1, 36, 2, 18, 3, 12, 4, 9, 6}) {
		t.Errorf("error: 0")
	}
}
