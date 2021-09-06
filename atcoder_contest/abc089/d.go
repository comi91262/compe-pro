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

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()
	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)
	var d int
	fmt.Fscan(reader, &d)

	a := make([][]int, h)
	b := make([]point, h*w+1)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &a[i][j])
			b[a[i][j]] = point{x: i, y: j}
		}
	}

	c := make([][]int, d+1)
	for i := 1; i <= d; i++ {
		for j := i; j < h*w+1-d; j += d {
			c[i] = append(c[i], abs(b[j+d].x-b[j].x)+abs(b[j+d].y-b[j].y))
		}
	}
	//fmt.Fprintf(writer, "%v\n", c)
	for i := 0; i < d+1; i++ {
		for j := 0; j < len(c[i])-1; j++ {
			c[i][j+1] += c[i][j]
		}
	}
	//fmt.Fprintf(writer, "%v\n", c)

	var q int
	fmt.Fscan(reader, &q)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		if r == l {
			fmt.Fprintf(writer, "%v\n", 0)
			continue
		}
		s := (l-1)%d + 1
		ans := c[s][(r-s)/d-1]
		if l-s != 0 {
			ans -= c[s][(l-s)/d-1]
		}
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
