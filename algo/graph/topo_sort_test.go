package graph

import (
	"testing"
)

func TestTopoSort(t *testing.T) {

	var g = make([][]int, 6)
	for i := 0; i < 6; i++ {
		g[i] = make([]int, 0)
		switch i {
		case 0:
			g[i] = append(g[i], 1)
			g[i] = append(g[i], 3)
		case 1:
			g[i] = append(g[i], 3)
		case 2:
			g[i] = append(g[i], 3)
			g[i] = append(g[i], 4)
		case 3:
		case 4:
			g[i] = append(g[i], 5)
		case 5:
		}
	}
	got := topoSort(g)

	want := []int{0, 2, 1, 4, 3, 5}

	if !testEqSlice(got, want) {
		t.Errorf("error: %v %v", got, want)
	}
}
