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
	parent = make([]int, n+1)
	rank = make([]int, n+1)

	for i := 0; i < n+1; i++ {
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

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	initUnionFind(200000)

	ans := 0
	for i := 0; i < n/2; i++ {
		x := a[i]
		y := a[n-i-1]

		if !same(x, y) {
			ans++
			unite(x, y)
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
	// fmt.Fprintf(writer, "%v\n", a)
	// fmt.Fprintf(writer, "%v\n", parent)
}
