package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func dfs(n, m, d, k int, a []int, acc *[][]int) {
	if n == d {
		b := make([]int, len(a))
		copy(b, a)
		*acc = append(*acc, b)
		return
	}

	for i := 0; k+i <= m; i++ {
		a = append(a, k+i)
		dfs(n, m, d+1, k+i, a, acc)
		a = a[:len(a)-1]
	}
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	var q int
	fmt.Fscan(reader, &q)
	var a = make([]int, q)
	var b = make([]int, q)
	var c = make([]int, q)
	var d = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		fmt.Fscan(reader, &c[i])
		fmt.Fscan(reader, &d[i])
		a[i]--
		b[i]--
	}

	ar := [][]int{}
	dfs(n, m, 0, 1, []int{}, &ar)

	ans := 0
	for _, x := range ar {
		cnt := 0
		for i := 0; i < q; i++ {
			if x[b[i]]-x[a[i]] == c[i] {
				cnt += d[i]
			}
		}
		ans = max(ans, cnt)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
