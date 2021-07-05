package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

func (r *RMQWithLazy) eval(k int) {
	if r.lazy[k] == 0 {
		return
	}
	if k < r.n-1 { // 葉でなければ子に伝搬
		r.lazy[k*2+1] = r.lazy[k]
		r.lazy[k*2+2] = r.lazy[k]
	}
	// 自身を更新
	r.data[k] = r.lazy[k]
	r.lazy[k] = 0
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
		rmq.data[k] = max(rmq.data[k*2+1], rmq.data[k*2+2])
	}
}

func (r *RMQWithLazy) query(a, b int) int { return r.querySub(a, b, 0, 0, r.n) }
func (rmq *RMQWithLazy) querySub(a, b, k, l, r int) int {
	rmq.eval(k)
	if r <= a || b <= l {
		return 0
	} else if a <= l && r <= b {
		return rmq.data[k]
	} else {
		vl := rmq.querySub(a, b, k*2+1, l, (l+r)/2)
		vr := rmq.querySub(a, b, k*2+2, (l+r)/2, r)
		return max(vl, vr)
	}
}

func main() {
	defer writer.Flush()

	var w, n int
	fmt.Fscan(reader, &w, &n)

	var l = make([]int, n)
	var r = make([]int, n)

	rmq := &RMQWithLazy{}
	rmq.create(w)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i], &r[i])
		h := rmq.query(l[i], r[i]+1) + 1
		rmq.update(l[i], r[i]+1, h)

		fmt.Fprintf(writer, "%d\n", h)
	}

}

// var n, q int
// fmt.Fscan(reader, &n, &q)

// var s string
// fmt.Fscan(reader, &n, &k, &s)
// ss := strings.Split(s, "")

// var a = make([]int, n)
// for i := 0; i < n; i++ {
// 	fmt.Fscan(reader, &a[i])
// }
// var n int
// fmt.Fscan(reader, &n)
// fmt.Fprintf(writer, "%d\n", n)

//	var n int
//	fmt.Fscan(reader, &n)
//	var a = make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Fscan(reader, &a[i])
//	}
// const d = 1_000_000_000 + 7
// for i := 0; i < n; i++ {
// for j := 0; j < n; j++ {
// for k := 0; k < n; k++ {
// }
// }
// }

// var g = make([][]int, n+1)
// var a = make([]int, m+1)
// var b = make([]int, m+1)
// for i := 1; i < m+1; i++ {
// 	fmt.Fscan(reader, &a[i], &b[i])
// 	g[a[i]] = append(g[a[i]], b[i])
// 	g[b[i]] = append(g[b[i]], a[i])
// }

// m := map[int]int{}
// a := make([]int, n)
// for i := 0; i < len(a); i++ {
// 	m[a[i]]++
// }

// func abs(x int) int {
// 	if x >= 0 {
// 		return x
// 	} else {
// 		return x * -1
// 	}
// }
//
// func min(arg ...int) int {
// 	min := arg[0]
// 	for _, x := range arg {
// 		if min > x {
// 			min = x
// 		}
// 	}
// 	return min
// }
//
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

/* Range Minumam Query：[0,n-1] について、区間ごとの最小値を管理する構造体
   update(a,b,x): 区間[a,b) の要素を x に更新。O(log(n))
   query(a,b): [a,b) での最小の要素を取得。O(log(n))
*/
