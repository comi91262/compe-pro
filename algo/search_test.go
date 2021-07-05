package algo

import "testing"

// 範囲外（左端) 気になるなら期待値は -1 であるように実装を変えてもいいかも
func TestBinerySearch1(t *testing.T) {
	a := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

	got := binarySearch(a, 0, lowerBound)
	want := 0
	if got != want {
		t.Errorf("binarySearch() = %d, want %d", got, want)
	}
}

// 範囲外（右端)
func TestBinerySearch2(t *testing.T) {
	a := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}
	got := binarySearch(a, 911, lowerBound)
	want := 10
	if got != want {
		t.Errorf("binarySearch() = %d, want %d", got, want)
	}
}

// 13が以上が条件を満たす。最小のindexは a[1] = 14 なので 1
func TestBinerySearch3(t *testing.T) {
	a := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}
	got := binarySearch(a, 13, lowerBound)
	want := 1
	if got != want {
		t.Errorf("binarySearch() = %d, want %d", got, want)
	}
}

// 14が以上が条件を満たす。最小のindexは a[1] = 14 なので 1
func TestBinerySearch4(t *testing.T) {
	a := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}
	got := binarySearch(a, 14, lowerBound)
	want := 1
	if got != want {
		t.Errorf("binarySearch() = %d, want %d", got, want)
	}
}

// 15が以上が条件を満たす。最小のindexは a[2] = 32 なので 2
func TestBinerySearch5(t *testing.T) {
	a := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}
	got := binarySearch(a, 15, lowerBound)
	want := 2
	if got != want {
		t.Errorf("binarySearch() = %d, want %d", got, want)
	}
}

// 51が以上が条件を満たす。a[x] = 51 な xの中で、最小のindexは 3
func TestBinerySearch6(t *testing.T) {
	a := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}
	got := binarySearch(a, 51, lowerBound)
	want := 3
	if got != want {
		t.Errorf("binarySearch() = %d, want %d", got, want)
	}

}
