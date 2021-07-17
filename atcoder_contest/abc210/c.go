package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type queue []int

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
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

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	q := queue{}
	m := map[int]int{}
	ans := 0
	cnt := 0
	for _, e := range c {
		q.push(e)
		if m[e] == 0 {
			cnt++
		}
		m[e]++

		for !q.empty() && len(q) > k {
			e := q.pop()

			if m[e] == 1 {
				cnt--
			}
			m[e]--
		}

		ans = max(ans, cnt)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
