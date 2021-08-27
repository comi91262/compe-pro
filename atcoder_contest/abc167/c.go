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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var x int
	fmt.Fscan(reader, &x)

	c := make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	ans := 1 << 60
	for i := 0; i < 1<<n; i++ {
		check := make([]int, m)
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				for k := 0; k < m; k++ {
					check[k] += a[j][k]
				}
			}
		}

		flag := true
		for j := 0; j < m; j++ {
			if check[j] < x {
				flag = false
			}
		}

		if flag {
			mn := 0
			for j := 0; j < n; j++ {
				if i&(1<<j) != 0 {
					mn += c[j]
				}
			}
			ans = min(mn, ans)
		}
	}
	if ans == 1<<60 {
		fmt.Fprintf(writer, "%v\n", -1)
	} else {
		fmt.Fprintf(writer, "%v\n", ans)
	}
}
