package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type UnionFind struct {
	parent []int
	rank   []int
}

func (uf *UnionFind) Create(n int) {
	uf.parent = make([]int, n+1)
	uf.rank = make([]int, n+1)

	for i := 0; i < n+1; i++ {
		uf.parent[i] = i
		uf.rank[i] = 0
	}
}

func (uf *UnionFind) root(x int) int {
	if uf.parent[x] == x {
		return x
	} else {
		uf.parent[x] = uf.root(uf.parent[x])
		return uf.parent[x]
	}
}

func (uf *UnionFind) IsSameRoot(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf *UnionFind) UniteTree(cx, cy int) {
	x := uf.root(cx)
	y := uf.root(cy)

	if x == y {
		return
	}

	if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	} else if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else {
		uf.parent[y] = x
		uf.rank[x]++
	}
}

type edge struct {
	from int
	to   int
	cost int
}

func kruskal(e []edge, n int) (int, int) {
	// コストが小さい順にソート
	sort.Slice(e, func(i, j int) bool { return e[i].cost < e[j].cost })

	uf := UnionFind{}
	uf.Create(n + 1)

	minCost := 0
	u := 0

	for i := 0; i < len(e); i++ {
		if !uf.IsSameRoot(e[i].from, e[i].to) {
			minCost += e[i].cost
			u++
			uf.UniteTree(e[i].from, e[i].to)
		}
	}

	return minCost, u
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	c := make([]int, m)
	l := make([]int, m)
	r := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &c[i], &l[i], &r[i])
	}

	es := []edge{}
	for i := 0; i < m; i++ {
		es = append(es, edge{from: l[i] - 1, to: r[i], cost: c[i]})
	}

	cost, u := kruskal(es, n)

	if u == n {
		fmt.Fprintf(writer, "%v\n", cost)
	} else {
		fmt.Fprintf(writer, "%v\n", -1)
	}

}
