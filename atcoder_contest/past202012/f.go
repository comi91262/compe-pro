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
	var c = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i], &c[i])
	}

	ans := 0
	for i := 0; i < 1<<n; i++ {
		ma := map[int]int{}
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				ma[j+1]++
			}
		}
		m2 := map[int]int{}
		for j := 0; j < m; j++ {
			cnt := 0
			dan := 0
			if ma[a[j]] > 0 {
				cnt++
			} else {
				dan = a[j]
			}
			if ma[b[j]] > 0 {
				cnt++
			} else {
				dan = b[j]

			}
			if ma[c[j]] > 0 {
				cnt++
			} else {
				dan = c[j]
			}

			if cnt == 2 {
				m2[dan]++
			}
		}
		ans = max(ans, len(m2))
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
