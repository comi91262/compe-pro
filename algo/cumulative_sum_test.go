package algo

import (
	"testing"
)

func TestCumulativeSum0(t *testing.T) {

	c := cumulativeSum2{}

	c.create(3, 2)

	var a = [][]int{
		{2, 4},
		{4, 6},
		{1, 9},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			c.add(i, j, a[i][j])
		}
	}

	c.build()

	got := c.get(0, 0, 3, 2)
	want := 26

	if got != want {
		t.Errorf("Error 0: %v", got)
	}
}
