package algo

import "sort"

type edge struct {
	from int
	to   int
	cost int
}

// 最小全域木のコストを求める (クラスカル法)
// n はノードの数
func kruskal(e []edge, n int) int {
	// コストが小さい順にソート
	sort.Slice(e, func(i, j int) bool { return e[i].cost < e[j].cost })

	uf := UnionFind{}
	uf.Create(n)

	minCost := 0

	for i := 0; i < len(e); i++ {
		if !uf.IsSameRoot(e[i].from, e[i].to) {
			minCost += e[i].cost
			uf.UniteTree(e[i].from, e[i].to)
		}
	}

	return minCost
}
