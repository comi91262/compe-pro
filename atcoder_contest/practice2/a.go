package main

import (
	"bufio"
	"bytes"
	"os"
)

var parent [200000]int

func root(x int) int {
	if parent[x] < 0 {
		return x
	} else {
		parent[x] = root(parent[x])
		return parent[x]
	}
}

func IsSameRoot(x, y int) bool {
	return root(x) == root(y)
}

func UniteTree(cx, cy int) {
	x := root(cx)
	y := root(cy)
	if x == y {
		return
	}
	if parent[x] > parent[y] {
		x, y = y, x
	}
	parent[x] += parent[y]
	parent[y] = x
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var buf bytes.Buffer
	buf.ReadFrom(os.Stdin)
	bytes := buf.Bytes()

	l := 0
	Int := func() (result int) {
		for bytes[l] < '0' || bytes[l] > '9' {
			l++
		}
		for '0' <= bytes[l] && bytes[l] <= '9' {
			result = result*10 + int(bytes[l]-'0')
			l++
		}
		return result
	}
	n, q := Int(), Int()
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	for i := 0; i < q; i++ {
		t, u, v := Int(), Int(), Int()
		switch t {
		case 0:
			UniteTree(u, v)
		case 1:
			if IsSameRoot(u, v) {
				writer.WriteByte('1')
			} else {
				writer.WriteByte('0')
			}
			writer.WriteByte('\n')
		}
	}
}
