package algo

// Library Check 検証済
// RMQ(Range Min Query): [0, n) について、最小値を管理する構造体
type RMQ struct {
	n    int
	data []int
}

func (r *RMQ) Create(n int) {
	r.data = make([]int, n*4)
	var x int = 1
	for n > x {
		x *= 2
	}
	r.n = x
}

// i 番目の要素をxに更新。O(log(n))
func (r *RMQ) Update(i, x int) {
	i += r.n - 1
	r.data[i] = x
	for i > 0 {
		i = (i - 1) / 2 // parent
		r.data[i] = min(r.data[i*2+1], r.data[i*2+2])
	}
}

// [a,b) での最小の要素を取得。O(log(n))
func (r *RMQ) Query(a, b int) int { return r.querySub(a, b, 0, 0, r.n) }
func (rmq *RMQ) querySub(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return 1 << 60
	} else if a <= l && r <= b {
		return rmq.data[k]
	} else {
		vl := rmq.querySub(a, b, k*2+1, l, (l+r)/2)
		vr := rmq.querySub(a, b, k*2+2, (l+r)/2, r)
		return min(vl, vr)
	}
}
