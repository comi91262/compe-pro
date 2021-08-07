package algo

/* Range Minumam Query：[0,n-1] について、区間ごとの最小値を管理する構造体
   update(a,b,x): 区間[a,b) の要素を x に更新。O(log(n))
   query(a,b): [a,b) での最小の要素を取得。O(log(n))
*/

import "math"

type RMQWithLazy struct {
	n    int
	data []int
	lazy []int
}

func (r *RMQWithLazy) create(n int) {
	r.data = make([]int, n*4)
	r.lazy = make([]int, n*4)
	var x int = 1
	for n > x {
		x *= 2
	}
	r.n = x
}

func (r *RMQWithLazy) min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func (r *RMQWithLazy) set(i, x int) { r.data[i+r.n-1] = x }

func (r *RMQWithLazy) build() {
	for k := r.n - 2; k >= 0; k-- {
		r.data[k] = min(r.data[2*k+1], r.data[2*k+2])
	}
}

func (r *RMQWithLazy) eval(k int) {
	if r.lazy[k] == math.MaxInt64 {
		return
	}
	if k < r.n-1 { // 葉でなければ子に伝搬
		r.lazy[k*2+1] = r.lazy[k]
		r.lazy[k*2+2] = r.lazy[k]
	}
	// 自身を更新
	r.data[k] = r.lazy[k]
	r.lazy[k] = math.MaxInt64
}

func (r *RMQWithLazy) update(a, b, x int) { r.updateSub(a, b, x, 0, 0, r.n) }
func (rmq *RMQWithLazy) updateSub(a, b, x, k, l, r int) {
	rmq.eval(k)
	if a <= l && r <= b { // 完全に内側の時
		rmq.lazy[k] = x
		rmq.eval(k)
	} else if a < r && l < b { // 一部区間が被る時
		rmq.updateSub(a, b, x, k*2+1, l, (l+r)/2) // 左の子
		rmq.updateSub(a, b, x, k*2+2, (l+r)/2, r) // 右の子
		rmq.data[k] = min(rmq.data[k*2+1], rmq.data[k*2+2])
	}
}

func (r *RMQWithLazy) query(a, b int) int { return r.querySub(a, b, 0, 0, r.n) }
func (rmq *RMQWithLazy) querySub(a, b, k, l, r int) int {
	rmq.eval(k)
	if r <= a || b <= l {
		return math.MaxInt64
	} else if a <= l && r <= b {
		return rmq.data[k]
	} else {
		vl := rmq.querySub(a, b, k*2+1, l, (l+r)/2)
		vr := rmq.querySub(a, b, k*2+2, (l+r)/2, r)
		return min(vl, vr)
	}
}
