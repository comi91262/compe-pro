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
	size   []int
}

func (uf *UnionFind) Create(n int) {
	uf.parent = make([]int, n)
	uf.size = make([]int, n)

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
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

	if uf.size[x] > uf.size[y] {
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
	} else {
		uf.parent[x] = y
		uf.size[y] += uf.size[x]
	}
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.root(x)]
}

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
			if e[i].cost > 0 {
				minCost += e[i].cost
			}
			uf.UniteTree(e[i].from, e[i].to)
		}
	}

	return minCost
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var e = make([]edge, m)
	var a = make([]int, m)
	var b = make([]int, m)
	var c = make([]int, m)

	sum := 0
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		fmt.Fscan(reader, &c[i])
		a[i]--
		b[i]--
		e[i].from = a[i]
		e[i].to = b[i]
		e[i].cost = c[i]
		if e[i].cost >= 0 {
			sum += e[i].cost
		}
	}

	cost := kruskal(e, n)
	fmt.Fprintf(writer, "%v\n", sum-cost)
}
