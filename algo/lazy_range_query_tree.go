package algo

import "math"

/* Lazy Range Query Tree：[0,n-1] について
   update(a,b,x): 区間[a,b) の要素を x に更新。O(log(n))
   query(a,b): [a,b) での最小の要素を取得。O(log(n))
*/
type LRQT struct {
	n    int
	data []int
	lazy []int
}

func (tree *LRQT) Create(n int) {
	tree.data = make([]int, n*4)
	tree.lazy = make([]int, n*4)
	var x int = 1
	for n > x {
		x *= 2
	}
	tree.n = x
}

func (tree *LRQT) eval(k int) {
	if tree.lazy[k] == math.MaxInt64 {
		return
	}
	if k < tree.n-1 { // 葉でなければ子に伝搬
		tree.lazy[k*2+1] = tree.lazy[k]
		tree.lazy[k*2+2] = tree.lazy[k]
	}
	tree.data[k] += tree.lazy[k]
	tree.lazy[k] = math.MaxInt64
}

func (tree *LRQT) Update(a, b, x int) { tree.updateSub(a, b, x, 0, 0, tree.n) }
func (tree *LRQT) updateSub(a, b, x, k, l, r int) {
	tree.eval(k)
	if a <= l && r <= b { // 完全に内側の時
		tree.lazy[k] = x
		tree.eval(k)
	} else if a < r && l < b { // 一部区間が被る時
		tree.updateSub(a, b, x, k*2+1, l, (l+r)/2) // 左の子
		tree.updateSub(a, b, x, k*2+2, (l+r)/2, r) // 右の子
		tree.data[k] = min(tree.data[k*2+1], tree.data[k*2+2])
	}
}

func (tree *LRQT) Query(a, b int) int { return tree.querySub(a, b, 0, 0, tree.n) }
func (tree *LRQT) querySub(a, b, k, l, r int) int {
	tree.eval(k)
	if r <= a || b <= l {
		return math.MaxInt64
	} else if a <= l && r <= b {
		return tree.data[k]
	} else {
		vl := tree.querySub(a, b, k*2+1, l, (l+r)/2)
		vr := tree.querySub(a, b, k*2+2, (l+r)/2, r)
		return min(vl, vr)
	}
}
