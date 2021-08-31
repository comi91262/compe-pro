package main

import (
	"bufio"
	"fmt"
	"os"
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

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	uf := UnionFind{}
	uf.Create(n)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		a--
		b--
		uf.UniteTree(a, b)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, uf.Size(i))
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
