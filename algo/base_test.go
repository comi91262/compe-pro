package algo

import "testing"

func TestByteToDigit(t *testing.T) {
	if i, _ := byteToDigit('0'); i != 0 {
		t.Errorf("error: %v", i)
	}
	if i, _ := byteToDigit('1'); i != 1 {
		t.Errorf("error: %v", i)
	}
	if i, _ := byteToDigit('8'); i != 8 {
		t.Errorf("error: %v", i)
	}
	if i, _ := byteToDigit('9'); i != 9 {
		t.Errorf("error: %v", i)
	}
}

func TestStrToDigits(t *testing.T) {
	if got := strToDigits(""); !testEqSlice(got, []int{}) {
		t.Errorf("error: got %v", got)
	}
	if got := strToDigits("123"); !testEqSlice(got, []int{1, 2, 3}) {
		t.Errorf("error: got %v", got)
	}
}

func TestToDigits(t *testing.T) {
	if got := toDigits(10, 10); !testEqSlice(got, []int{1, 0}) {
		t.Errorf("error: got %v", got)
	}
	if got := toDigits(10, 2); !testEqSlice(got, []int{1, 0, 1, 0}) {
		t.Errorf("error: got %v", got)
	}
}

func TestToNumber(t *testing.T) {
	if got := toNumber([]int{1, 0, 1, 0}, 2); got != 10 {
		t.Errorf("error: got %v", got)
	}
	if got := toNumber([]int{1, 0}, 10); got != 10 {
		t.Errorf("error: got %v", got)
	}
}

func TestTo2BasedDisits(t *testing.T) {
	if got := toBinaryDigits(10, 8); !testEqSlice(got, []int{0, 0, 0, 0, 1, 0, 1, 0}) {
		t.Errorf("error: got %v", got)
	}
	if got := toBinaryDigits(0, 2); !testEqSlice(got, []int{0, 0}) {
		t.Errorf("error: got %v", got)
	}
}
