package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var x int
	fmt.Fscan(reader, &x)
	var m int
	fmt.Fscan(reader, &m)

	used := make([]int, m)
	used[x] = 1
	start := 0
	ans := x
	pre := x
	for i := 2; i <= n; i++ {
		nex := pre * pre % m
		if used[nex] != 0 {
			start = i
			break
		}
		ans += nex
		used[nex] = i
		pre = nex
	}

	n -= start
	if start == 0 {
		fmt.Fprintf(writer, "%v\n", ans)
		return
	}

	cycle := []int{}
	used = make([]int, m)
	for i := 1; i <= n; i++ {
		nex := pre * pre % m
		if used[nex] != 0 {
			break
		}
		cycle = append(cycle, nex)
		used[nex] = i
		pre = nex
	}

	for i := 0; i < len(cycle)-1; i++ {
		cycle[i+1] += cycle[i]
	}

	ans += n / len(cycle) * cycle[len(cycle)-1]
	n %= len(cycle)
	ans += cycle[n]

	fmt.Fprintf(writer, "%v\n", ans)
}
