package graph

type LCA struct {
	parent [][]int // parent[k][u]:= u の 2^k 先の親
	depth  []int
	dist   []int // root からの距離
}

func (lca *LCA) Create(root int, G [][]edge) {
	V := len(G)
	K := 1
	for (1 << K) < V {
		K++
	}
	lca.parent = make([][]int, K)
	for i := 0; i < K; i++ {
		lca.parent[i] = make([]int, V)
		for j := 0; j < V; j++ {
			lca.parent[i][j] = -1
		}
	}
	lca.depth = make([]int, V)
	lca.dist = make([]int, V)
	for i := 0; i < V; i++ {
		lca.depth[i] = -1
		lca.dist[i] = -1
	}

	lca.dfs(G, root, -1, 0, 0)
	for k := 0; k+1 < K; k++ {
		for v := 0; v < V; v++ {
			if lca.parent[k][v] < 0 {
				lca.parent[k+1][v] = -1
			} else {
				lca.parent[k+1][v] = lca.parent[k][lca.parent[k][v]]
			}
		}
	}
}

// 根からの距離と1つ先の頂点を求める
func (lca *LCA) dfs(G [][]edge, v, p, de, di int) {
	lca.parent[0][v] = p
	lca.depth[v] = de
	lca.dist[v] = di
	for _, e := range G[v] {
		if e.to != p {
			lca.dfs(G, e.to, v, de+1, di+e.cap)
		}
	}
}

func (lca *LCA) Query(u, v int) int {
	if lca.depth[u] < lca.depth[v] {
		u, v = v, u
	}
	K := len(lca.parent)
	// LCA までの距離を同じにする
	for k := 0; k < K; k++ {
		if ((lca.depth[u]-lca.depth[v])>>k)&1 != 0 {
			u = lca.parent[k][u]
		}
	}
	// 二分探索で LCA を求める
	if u == v {
		return u
	}
	for k := K - 1; k >= 0; k-- {
		if lca.parent[k][u] != lca.parent[k][v] {
			u = lca.parent[k][u]
			v = lca.parent[k][v]
		}
	}
	return lca.parent[0][u]
}

func (lca *LCA) Dist(u, v int) int {
	return lca.dist[u] + lca.dist[v] - 2*lca.dist[lca.Query(u, v)]
}
