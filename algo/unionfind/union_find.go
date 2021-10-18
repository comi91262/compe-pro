package unionfind

type UnionFind struct {
	parent []int
	size   []int
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

func New(n int) *UnionFind {
	uf := UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return &uf
}
