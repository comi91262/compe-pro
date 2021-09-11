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

func (q *queue) front() int {
	return (*q)[0]
}

func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	a := make([]queue, m)
	cnt := make([][]int, n)
	for i := 0; i < m; i++ {
		var k int
		fmt.Fscan(reader, &k)

		a[i] = queue{}
		for j := 0; j < k; j++ {
			var tmp int
			fmt.Fscan(reader, &tmp)
			a[i].push(tmp - 1)
		}
		cnt[a[i].front()] = append(cnt[a[i].front()], i)
	}

	q := queue{}
	for i := 0; i < n; i++ {
		if len(cnt[i]) == 2 {
			q.push(i)
		}
	}

	for !q.empty() {
		v := q.pop()
		for _, c := range cnt[v] {
			a[c].pop()
			if !a[c].empty() {
				cnt[a[c].front()] = append(cnt[a[c].front()], c)
				if len(cnt[a[c].front()]) == 2 {
					q.push(a[c].front())
				}
			}
		}
	}

	for i := range a {
		if len(a[i]) != 0 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
