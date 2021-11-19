package algo

import (
	"strconv"
)

type RQT struct {
	n    int
	x    []int
	unit int
	op   func(x ...int) int
}

func NewRangeQueryTree(seq []int, op func(x ...int) int, unit int) *RQT {
	tree := RQT{}

	tree.n = len(seq)
	tree.x = make([]int, len(seq)*2)
	tree.unit = unit
	tree.op = op

	for i := range tree.x {
		tree.x[i] = tree.unit
	}
	for i, x := range seq {
		tree.x[i+len(seq)] = x
	}
	for i := len(seq) - 1; i > 0; i-- {
		tree.x[i] = tree.op(tree.x[i<<1], tree.x[i<<1|1])
	}

	return &tree
}

func (tree *RQT) Update(i, x int) {
	i += tree.n
	tree.x[i] = x
	for i > 1 {
		i >>= 1
		tree.x[i] = tree.op(tree.x[i<<1], tree.x[i<<1|1])
	}
}

func (tree *RQT) Query(l, r int) int {
	l += tree.n
	r += tree.n
	vl := tree.unit
	vr := tree.unit

	for l < r {
		if l&1 > 0 {
			vl = tree.op(vl, tree.x[l])
			l += 1
		}
		if r&1 > 0 {
			r -= 1
			vr = tree.op(tree.x[r], vr)
		}
		l >>= 1
		r >>= 1
	}
	return tree.op(vl, vr)
}

func (tree *RQT) Debug() (ret string) {
	ret = ""
	for i := 0; i < tree.n; i++ {
		if i > 0 {
			ret += " "
		}
		ret += strconv.Itoa(tree.x[i+tree.n])
	}
	return
}
