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

func main() {
	defer writer.Flush()

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	q := queue{}

	for i := 0; i < n; i++ {
		for j := 0; j < a[i]; j++ {
			q.push(i + 1)
		}
	}
	ans := make([][]int, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans[i][j] = q.pop()
		}
	}
	for i := 0; i < h; i++ {
		if i%2 == 1 {
			for j := 0; j < w/2; j++ {
				ans[i][j], ans[i][w-j-1] = ans[i][w-j-1], ans[i][j]
			}
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprintf(writer, "%v", ans[i][j])
			if j != w-1 {
				fmt.Fprintf(writer, " ")
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}
