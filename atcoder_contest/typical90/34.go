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

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	q := queue{}
	m := map[int]int{}

	cnt := 0
	ans := 0
	for _, e := range a {
		q.push(e)
		if m[e] == 0 {
			cnt++
		}
		m[e]++

		for !q.empty() && cnt > k {
			e := q.pop()

			if m[e] == 1 {
				cnt--
			}
			m[e]--
		}

		ans = max(ans, len(q))
	}
	//	ans := 0
	//	cr := 0
	//	cnt := 0
	//	m := map[int]int{}
	//	for i := 0; i < n; i++ {
	//		for cr < n {
	//			if m[a[cr]] == 0 && cnt == k {
	//				break
	//			}
	//			if m[a[cr]] == 0 {
	//				cnt++
	//			}
	//			m[a[cr]] += 1
	//			cr++
	//		}
	//		ans = max(ans, cr-i)
	//		if m[a[i]] == 1 {
	//			cnt--
	//		}
	//		m[a[i]] -= 1
	//	}

	fmt.Fprintf(writer, "%v\n", ans)
}
