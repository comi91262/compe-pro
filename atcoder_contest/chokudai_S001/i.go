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

func (q *queue) front() int {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	q := queue{}
	ans := 0
	sum := 0
	for i := 0; i < n; i++ {
		q.push(a[i])
		sum += a[i]
		for !q.empty() && sum > n {
			sum -= q.pop()
		}
		if sum == n {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
