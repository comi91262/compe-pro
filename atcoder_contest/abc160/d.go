package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	m := map[int]int{}
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if i == j {
				continue
			}

			cand := abs(i - j)
			cand = min(cand, 1+abs(x-i)+abs(y-j))
			m[cand]++
		}
	}
	for i := 1; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", m[i])
	}
}
