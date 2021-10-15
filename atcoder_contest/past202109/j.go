package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}

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
	if tree.lazy[k] == 0 {
		return
	}
	if k < tree.n-1 { // 葉でなければ子に伝搬
		tree.lazy[k*2+1] += tree.lazy[k]
		tree.lazy[k*2+2] += tree.lazy[k]
	}
	tree.data[k] += tree.lazy[k]
	tree.lazy[k] = 0
}

func (tree *LRQT) Update(a, b, x int) { tree.updateSub(a, b, x, 0, 0, tree.n) }
func (tree *LRQT) updateSub(a, b, x, k, l, r int) {
	tree.eval(k)
	if a <= l && r <= b { // 完全に内側の時
		tree.lazy[k] += x
		tree.eval(k)
	} else if a < r && l < b { // 一部区間が被る時
		tree.updateSub(a, b, x, k*2+1, l, (l+r)/2) // 左の子
		tree.updateSub(a, b, x, k*2+2, (l+r)/2, r) // 右の子
		tree.data[k] = add(tree.data[k*2+1], tree.data[k*2+2])
	}
}

func (tree *LRQT) Query(a, b int) int { return tree.querySub(a, b, 0, 0, tree.n) }
func (tree *LRQT) querySub(a, b, k, l, r int) int {
	tree.eval(k)
	if r <= a || b <= l {
		return 0
	} else if a <= l && r <= b {
		return tree.data[k]
	} else {
		vl := tree.querySub(a, b, k*2+1, l, (l+r)/2)
		vr := tree.querySub(a, b, k*2+2, (l+r)/2, r)
		return add(vl, vr)
	}
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var q int
	fmt.Fscan(reader, &q)

	a := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		a[i] = i + 1
	}

	tree := LRQT{}
	tree.Create(2 * n)

	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		var k int
		fmt.Fscan(reader, &k)

		switch t {
		case 1:
			k--
			cnt := tree.Query(k, k+1)
			//			for i := 0; i < 2*n; i++ {
			//				fmt.Fprintf(writer, "%v ", tree.Query(i, i+1))
			//			}
			//			fmt.Fprintf(writer, "\n")
			if cnt%2 == 0 {
				fmt.Fprintf(writer, "%v\n", a[k])
			} else {
				fmt.Fprintf(writer, "%v\n", a[n*2-k-1])
			}
		case 2:
			tree.Update(n-k, n+k, 1)
		}
	}
}
