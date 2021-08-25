package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var a = make([]int, m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
	}
	var k int
	fmt.Fscan(reader, &k)
	var c = make([]int, k)
	var d = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &c[i])
		fmt.Fscan(reader, &d[i])
	}

	ans := 0
	for i := 0; i < 1<<k; i++ {
		ma := map[int]int{}
		for j := 0; j < k; j++ {
			if i&(1<<j) != 0 {
				ma[c[j]]++
			} else {
				ma[d[j]]++
			}
		}

		cnt := 0
		for j := 0; j < m; j++ {
			if ma[a[j]] > 0 && ma[b[j]] > 0 {
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
