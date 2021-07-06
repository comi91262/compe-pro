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

	for i := 1; i < n+1; i++ {
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

	var t = make([]int, q)
	var x = make([]int, q)
	var y = make([]int, q)
	var v = make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &t[i], &x[i], &y[i], &v[i])
		x[i]--
		y[i]--
	}

	var s = make([]int, n-1)
	var p = make([]int, n)

	for i := 0; i < q; i++ {
		if t[i] == 0 {
			s[x[i]] = v[i]
		}
	}

	for i := 0; i < n-1; i++ {
		p[i+1] = s[i] - p[i]
	}

	initUnionFind(n)

	for i := 0; i < q; i++ {
		switch {
		case t[i] == 0:
			unite(x[i], y[i])
		case t[i] == 1:
			if !same(x[i], y[i]) {
				fmt.Fprintf(writer, "Ambiguous\n")
			} else if (x[i]+y[i])%2 == 0 {
				fmt.Fprintf(writer, "%v\n", p[y[i]]+v[i]-p[x[i]])
			} else {
				fmt.Fprintf(writer, "%v\n", p[y[i]]-(v[i]-p[x[i]]))
			}
		}
	}
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
