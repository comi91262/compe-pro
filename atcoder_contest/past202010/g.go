package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type UnionFind struct {
	parent []int
	size   []int
}

func (uf *UnionFind) Create(n int) {
	uf.parent = make([]int, n)
	uf.size = make([]int, n)

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
}

func (uf *UnionFind) root(x int) int {
	if uf.parent[x] == x {
		return x
	} else {
		uf.parent[x] = uf.root(uf.parent[x])
		return uf.parent[x]
	}
}

func (uf *UnionFind) IsSameRoot(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf *UnionFind) UniteTree(cx, cy int) {
	x := uf.root(cx)
	y := uf.root(cy)

	if x == y {
		return
	}

	if uf.size[x] > uf.size[y] {
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
	} else {
		uf.parent[x] = y
		uf.size[y] += uf.size[x]
	}
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.root(x)]
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	s := make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}
	uf := UnionFind{}
	uf.Create(n * m)

	dx := [4]int{0, 0, 1, -1}
	dy := [4]int{1, -1, 0, 0}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == "." {
				cnt++
			}
			for k := 0; k < 4; k++ {
				nx := i + dx[k]
				ny := j + dy[k]
				if 0 <= nx && nx < n && 0 <= ny && ny < m {
					if s[i][j] == "." && s[nx][ny] == "." {
						uf.UniteTree(i*m+j, nx*m+ny)
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == "#" {
				u := UnionFind{}
				u.Create(n * m)
				copy(u.parent, uf.parent)
				copy(u.size, uf.size)

				for k := 0; k < 4; k++ {
					nx := i + dx[k]
					ny := j + dy[k]
					if 0 <= nx && nx < n && 0 <= ny && ny < m {
						if s[nx][ny] == "." {
							u.UniteTree(i*m+j, nx*m+ny)
						}
					}
				}
				if u.Size(i*m+j) == cnt+1 {
					ans++
				}
			}

		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
