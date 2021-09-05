package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type point struct {
	x int
	y int
}

type queue []point

func (q *queue) push(n point) {
	*q = append(*q, n)
}

func (q *queue) pop() point {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func diff(p1, p2 point) point {
	return point{x: p1.x - p2.x, y: p1.y - p2.y}
}

func prod(p1, p2 point) int {
	return p1.x*p2.y - p1.y*p2.x
}

func main() {
	defer writer.Flush()
	var ax int
	fmt.Fscan(reader, &ax)
	var ay int
	fmt.Fscan(reader, &ay)
	var bx int
	fmt.Fscan(reader, &bx)
	var by int
	fmt.Fscan(reader, &by)

	var n int
	fmt.Fscan(reader, &n)

	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	q1 := queue{}
	q2 := queue{}

	for i := 0; i < n; i++ {
		if i == n-1 {
			q1.push(point{x[0], y[0]})
			q2.push(point{x[n-1], y[n-1]})
			continue
		}
		q1.push(point{x[i], y[i]})
		q2.push(point{x[i+1], y[i+1]})
	}

	cnt := 0
	pa := point{x: ax, y: ay}
	pb := point{x: bx, y: by}
	for !q1.empty() && !q2.empty() {
		p1 := q1.pop()
		p2 := q2.pop()

		s1 := prod(diff(pb, pa), diff(p1, pa)) * prod(diff(pb, pa), diff(p2, pa))
		s2 := prod(diff(p2, p1), diff(pa, p1)) * prod(diff(p2, p1), diff(pb, p1))

		if s1 <= 0 && s2 <= 0 {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt/2+1)
}
