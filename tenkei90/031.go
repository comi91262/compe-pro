package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func mex(a []int) int {
	sort.Ints(a)
	mx := 0
	for _, x := range a {
		if x < mx {
			continue
		}
		if x != mx {
			break
		}
		mx++
	}
	return mx
}

func grundy(w, b int) [][]int {
	g := make([][]int, w+1)
	for i := 0; i < w+1; i++ {
		g[i] = make([]int, b+w*(w+1)/2+1)
	}

	for i := 2; i < b+w*(w+1)/2+1; i++ {
		set := []int{}
		for j := 1; j <= i/2; j++ {
			set = append(set, g[0][i-j])
		}
		g[0][i] = mex(set)
	}

	for i := 1; i <= w; i++ {
		for j := 0; j < b+w*(w+1)/2+1-i; j++ {
			set := []int{}
			set = append(set, g[i-1][j+i])
			for k := 1; k <= j/2; k++ {
				set = append(set, g[i][j-k])
			}
			g[i][j] = mex(set)
		}
	}

	return g
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	g := grundy(50, 50)

	xor := 0
	for i := 0; i < n; i++ {
		xor ^= g[a[i]][b[i]]
	}

	if xor == 0 {
		fmt.Fprintf(writer, "%v\n", "Second")
	} else {
		fmt.Fprintf(writer, "%v\n", "First")
	}
}
