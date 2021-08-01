package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

// Library Checker 検証済
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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var g = make([][]int, n+1)
	var a = make([]int, m+1)
	var b = make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	ans := 0
	for i := 1; i < m+1; i++ {
		uf := UnionFind{}
		uf.Create(n)
		for j := 1; j < m+1; j++ {
			if i == j {
				continue
			}
			uf.UniteTree(a[j], b[j])
		}
		if !uf.IsSameRoot(a[i], b[i]) {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
