package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type edge struct {
	p1 point
	p2 point
}

type point struct {
	x int
	y int
}

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var p = make([]point, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(reader, &x)
		fmt.Fscan(reader, &y)
		p[i].x = x
		p[i].y = y
	}

	m := map[int][]int{}
	for i := 0; i < n; i++ {
		m[p[i].x] = append(m[p[i].x], p[i].y)
	}

	e := []edge{}
	for k, v := range m {
		if len(v) < 1 {
			continue
		}
		sort.Ints(v)
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				p1 := point{k, v[i]}
				p2 := point{k, v[j]}
				e = append(e, edge{p1, p2})
			}
		}
	}

	if len(e) == 0 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	ans := 0
	for i := 0; i < len(e); i++ {
		for j := i + 1; j < len(e); j++ {
			if e[i].p1.x == e[j].p1.x {
				continue
			}
			if e[i].p1.y == e[j].p1.y && e[i].p2.y == e[j].p2.y {
				ans++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
