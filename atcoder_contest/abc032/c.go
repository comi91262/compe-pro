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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)
	var s = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	for i := 0; i < n; i++ {
		if s[i] == 0 {
			fmt.Fprintf(writer, "%v\n", n)
			return
		}
	}

	ans := 0
	q := queue{}
	prod := 1
	for i := 0; i < n; i++ {
		q.push(s[i])
		prod *= s[i]
		for !q.empty() && prod > k {
			a := q.pop()
			prod /= a
		}
		ans = max(ans, len(q))
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
