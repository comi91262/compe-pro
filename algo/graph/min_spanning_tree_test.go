package algo

import "testing"

func TestKruskal(t *testing.T) {
	got := kruskal([]edge{
		{from: 0, to: 1, cost: 10},
		{from: 0, to: 3, cost: 5},
		{from: 1, to: 2, cost: 1},
		{from: 1, to: 3, cost: 1000},
		{from: 1, to: 4, cost: 500},
		{from: 2, to: 3, cost: 100},
		{from: 2, to: 4, cost: 10000},
		{from: 3, to: 4, cost: 5000},
	}, 5)
	want := 516
	if got != want {
		t.Errorf("error: %v", got)
	}
}
