package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func dfs(s, pre int, g [][]int) int {
	cnt := 0
	mn := 1 << 60
	mx := 0
	total := 0
	for _, nex := range g[s] {
		if nex == pre {
			continue
		}
		cnt++
		total = dfs(nex, s, g)
		mn = min(mn, total)
		mx = max(mx, total)
	}

	switch cnt {
	case 0:
		return 1
	case 1:
		return 2*total + 1
	default:
		return mn + mx + 1
	}
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var g = make([][]int, n)
	for i := 0; i < n-1; i++ {
		var a int
		fmt.Fscan(reader, &a)
		a--
		g[a] = append(g[a], i+1)
	}

	fmt.Fprintf(writer, "%v\n", dfs(0, 0, g))
}
