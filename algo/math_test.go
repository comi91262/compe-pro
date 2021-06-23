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
