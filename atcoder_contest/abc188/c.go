package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	x int
	y int
}

type queue []pair

func (q *queue) push(n pair) {
	*q = append(*q, n)
}

func (q *queue) pop() pair {
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

	q := queue{}

	var a = make([]pair, 1<<n)
	for i := 0; i < 1<<n; i++ {
		fmt.Fscan(reader, &a[i].x)
		a[i].y = i + 1
		q.push(a[i])
	}

	for len(q) > 2 {
		x := q.pop()
		y := q.pop()
		if x.x > y.x {
			q.push(x)
		} else {
			q.push(y)
		}
	}

	x := q.pop()
	y := q.pop()
	if x.x > y.x {
		fmt.Fprintf(writer, "%v\n", y.y)
	} else {
		fmt.Fprintf(writer, "%v\n", x.y)
	}

}
