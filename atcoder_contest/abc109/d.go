package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	y  int
	x  int
	yy int
	xx int
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
	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	q := queue{}
	for i := 0; i < h; i++ {
		if i%2 == 0 {
			for j := 0; j < w; j++ {
				if a[i][j]%2 == 0 {
					continue
				}
				if j+1 < w {
					q.push(pair{y: j + 1, x: i + 1, yy: j + 2, xx: i + 1})
					a[i][j]--
					a[i][j+1]++
				} else {
					if i+1 < h {
						q.push(pair{y: j + 1, x: i + 1, yy: j + 1, xx: i + 2})
						a[i][j]--
						a[i+1][j]++
					}
				}
			}
		} else {
			for j := w - 1; j >= 0; j-- {
				if a[i][j]%2 == 0 {
					continue
				}
				if j-1 >= 0 {
					q.push(pair{y: j + 1, x: i + 1, yy: j, xx: i + 1})
					a[i][j]--
					a[i][j-1]++
				} else {
					if i+1 < h {
						q.push(pair{y: j + 1, x: i + 1, yy: j + 1, xx: i + 2})
						a[i][j]--
						a[i+1][j]++
					}
				}
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", len(q))
	for !q.empty() {
		p := q.pop()
		fmt.Fprintf(writer, "%v %v %v %v\n", p.x, p.y, p.xx, p.yy)
	}
}
