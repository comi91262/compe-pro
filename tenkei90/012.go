package main

import (
	"bufio"
	"fmt"
	"os"
)

var parent []int
var rank []int

func initUnionFind(hw int) {
	parent = make([]int, hw)
	rank = make([]int, hw)
	for i := 0; i < hw; i++ {
		parent[i] = i
		rank[i] = 0
	}
}

func root(pos int) int {
	p := parent[pos]

	if p == pos {
		return p
	} else {
		gp := root(p)
		parent[pos] = gp
		return gp
	}
}

func same(pos1, pos2 int) bool {
	return root(pos1) == root(pos2)
}

func unite(pos1, pos2 int) {
	r1 := root(pos1)
	r2 := root(pos2)

	if r1 == r2 {
		return
	}

	if rank[pos1] < rank[pos2] {
		parent[pos1] = pos2
	} else if rank[pos1] > rank[pos2] {
		parent[pos2] = pos1
	} else if rank[pos1] == rank[pos2] {
		parent[pos2] = pos1
		rank[pos1]++
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var h, w, q int
	fmt.Fscan(reader, &h, &w, &q)

	var t = make([]int, q)
	var r = make([]int, q)
	var c = make([]int, q)
	var ra = make([]int, q)
	var ca = make([]int, q)
	var rb = make([]int, q)
	var cb = make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &t[i])
		switch t[i] {
		case 1:
			fmt.Fscan(reader, &r[i], &c[i])
		case 2:
			fmt.Fscan(reader, &ra[i], &ca[i], &rb[i], &cb[i])
		}
	}

	initUnionFind(h * w)

	var used [][]bool
	used = make([][]bool, h*w)
	for i := 0; i < h; i++ {
		used[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			used[i][j] = false
		}
	}

	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			var x, y = r[i] - 1, c[i] - 1
			used[x][y] = true
			area := parent[x*w+y]

			if x > 0 && used[x-1][y] {
				unite(area, parent[(x-1)*w+y])
				area = root(area)
			}

			if x+1 < h && used[x+1][y] {
				unite(area, parent[(x+1)*w+y])
				area = root(area)
			}

			if y > 0 && used[x][y-1] {
				unite(area, parent[x*w+y-1])
				area = root(area)
			}

			if y+1 < w && used[x][y+1] {
				unite(area, parent[x*w+y+1])
				area = root(area)
			}

		case 2:
			p1 := parent[(ra[i]-1)*w+ca[i]-1]
			p2 := parent[(rb[i]-1)*w+cb[i]-1]

			if used[ra[i]-1][ca[i]-1] && used[rb[i]-1][cb[i]-1] && same(p1, p2) {
				fmt.Fprintf(writer, "Yes\n")

			} else {
				fmt.Fprintf(writer, "No\n")
			}
		}
	}

}
