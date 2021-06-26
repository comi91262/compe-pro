package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var parent []int
var rank []int

func initUnionFind(n int) {
	parent = make([]int, n)
	rank = make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
}

func root(x int) int {
	if parent[x] == x {
		return x
	} else {
		parent[x] = root(parent[x])
		return parent[x]
	}
}

func same(x, y int) bool {
	return root(x) == root(y)
}

func unite(cx, cy int) {
	x := root(cx)
	y := root(cy)

	if x == y {
		return
	}

	if rank[x] < rank[y] {
		parent[x] = y
	} else if rank[x] > rank[y] {
		parent[y] = x
	} else {
		parent[y] = x
		rank[x]++
	}
}

func main() {
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	var p = make([]int, q)
	var a = make([]int, q)
	var b = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &p[i], &a[i], &b[i])
	}

	initUnionFind(n)

	for i := 0; i < q; i++ {
		if p[i] == 0 {
			unite(a[i], b[i])
		} else {
			if same(a[i], b[i]) {
				fmt.Fprintf(writer, "Yes\n")
			} else {
				fmt.Fprintf(writer, "No\n")
			}
		}
	}

}
