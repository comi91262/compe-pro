package algo

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
