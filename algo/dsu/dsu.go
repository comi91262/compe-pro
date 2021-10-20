package dsu

type DisjointSetForest struct {
	parent []int
	size   []int
}

func (dsu *DisjointSetForest) root(x int) int {
	if dsu.parent[x] == x {
		return x
	} else {
		dsu.parent[x] = dsu.root(dsu.parent[x])
		return dsu.parent[x]
	}
}

func (dsu *DisjointSetForest) IsSameRoot(x, y int) bool {
	return dsu.root(x) == dsu.root(y)
}

func (dsu *DisjointSetForest) UniteTree(cx, cy int) {
	x := dsu.root(cx)
	y := dsu.root(cy)

	if x == y {
		return
	}

	if dsu.size[x] > dsu.size[y] {
		dsu.parent[y] = x
		dsu.size[x] += dsu.size[y]
	} else {
		dsu.parent[x] = y
		dsu.size[y] += dsu.size[x]
	}
}

func (dsu *DisjointSetForest) Size(x int) int {
	return dsu.size[dsu.root(x)]
}

func NewDSU(n int) *DisjointSetForest {
	dsu := DisjointSetForest{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return &dsu
}
