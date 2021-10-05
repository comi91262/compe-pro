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

type edge struct {
	to   int
	cost int
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var k int
	fmt.Fscan(reader, &k)

	var g = make([][]edge, n)
	var a = make([]int, m)
	var b = make([]int, m)
	uf := UnionFind{}
	uf.Create(n)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], edge{b[i], 0})
		g[b[i]] = append(g[b[i]], edge{a[i], 0})
		uf.UniteTree(a[i], b[i])
	}

	bl := make([][]int, n)
	for i := 0; i < k; i++ {
		var c int
		fmt.Fscan(reader, &c)
		var d int
		fmt.Fscan(reader, &d)
		c--
		d--
		bl[c] = append(bl[c], d)
		bl[d] = append(bl[d], c)
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprintf(writer, " ")
		}
		ans := uf.Size(i) - len(g[i]) - 1
		for _, j := range bl[i] {
			if uf.IsSameRoot(i, j) {
				ans--
			}
		}
		fmt.Fprintf(writer, "%v", ans)
	}
	fmt.Fprintf(writer, "\n")
}
