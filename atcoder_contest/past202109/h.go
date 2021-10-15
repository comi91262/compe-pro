package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var parent [][]int // parent[k][u]:= u の 2^k 先の親
var depth []int
var dist []int // root からの距離

// 初期化
func create(root int, G [][]edge) {
	V := len(G)
	K := 1
	for (1 << K) < V {
		K++
	}
	parent = make([][]int, K)
	for i := 0; i < K; i++ {
		parent[i] = make([]int, V)
		for j := 0; j < V; j++ {
			parent[i][j] = -1
		}
	}
	depth = make([]int, V)
	dist = make([]int, V)
	for i := 0; i < V; i++ {
		depth[i] = -1
		dist[i] = -1
	}

	dfs(G, root, -1, 0, 0)
	for k := 0; k+1 < K; k++ {
		for v := 0; v < V; v++ {
			if parent[k][v] < 0 {
				parent[k+1][v] = -1
			} else {
				parent[k+1][v] = parent[k][parent[k][v]]
			}
		}
	}
}

// 根からの距離と1つ先の頂点を求める
func dfs(G [][]edge, v, p, de, di int) {
	parent[0][v] = p
	depth[v] = de
	dist[v] = di
	for _, e := range G[v] {
		if e.to != p {
			dfs(G, e.to, v, de+1, di+e.cap)
		}
	}
}

func query(u, v int) int {
	if depth[u] < depth[v] {
		u, v = v, u
	}
	K := len(parent)
	// LCA までの距離を同じにする
	for k := 0; k < K; k++ {
		if ((depth[u]-depth[v])>>k)&1 != 0 {
			u = parent[k][u]
		}
	}
	// 二分探索で LCA を求める
	if u == v {
		return u
	}
	for k := K - 1; k >= 0; k-- {
		if parent[k][u] != parent[k][v] {
			u = parent[k][u]
			v = parent[k][v]
		}
	}
	return parent[0][u]
}

type edge struct {
	to  int
	cap int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var x int
	fmt.Fscan(reader, &x)

	var g = make([][]edge, n)
	var a = make([]int, n-1)
	var b = make([]int, n-1)
	var c = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &a[i], &b[i], &c[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], edge{b[i], c[i]})
		g[b[i]] = append(g[b[i]], edge{a[i], c[i]})
	}

	create(0, g)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := dist[i] + dist[j] - 2*dist[query(i, j)]
			if d == x {
				fmt.Fprintf(writer, "%v\n", "Yes")
				return
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")
}
