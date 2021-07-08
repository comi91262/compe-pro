package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	m := 5001
	c := make([][]int, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		c[a[i]][b[i]]++
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m-1; j++ {
			c[i][j+1] += c[i][j]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m-1; j++ {
			c[j+1][i] += c[j][i]
		}
	}

	mx := 0
	for a := 1; a <= m-k-1; a++ {
		for b := 1; b <= m-k-1; b++ {
			mx = max(mx, c[a+k][b+k]-c[a-1][b+k]-c[a+k][b-1]+c[a-1][b-1])
		}
	}

	fmt.Fprintf(writer, "%v\n", mx)
}
