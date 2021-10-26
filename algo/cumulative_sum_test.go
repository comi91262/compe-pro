package algo

import (
	"testing"
)

func TestCumulativeSum0(t *testing.T) {

	c := NewPrefixSum(3, 2)

	var a = [][]int{
		{2, 4},
		{4, 6},
		{1, 9},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			c.Add(i, j, a[i][j])
		}
	}

	c.Build()
	got := c.Get(-1, -1, 2, 1)
	want := 26

	if got != want {
		t.Errorf("Error 0: %v", got)
	}
}
