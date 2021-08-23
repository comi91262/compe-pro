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
	var m int
	fmt.Fscan(reader, &m)

	a := make([][]int, m)
	k := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)

		fmt.Fscan(reader, &k[i])

		for j := 0; j < k[i]; j++ {
			fmt.Fscan(reader, &a[i][j])
			a[i][j]--
		}
		for j := k[i]; j < n; j++ {
			a[i][j]--
		}
	}

	var p = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &p[i])
	}

	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, m)
		for j := 0; j < m; j++ {
			for _, k := range a[j] {
				if k == i {
					b[i][j]++
				}
			}
		}
	}

	ans := 0
	for i := 0; i < 1<<n; i++ {
		c := make([]int, m)
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				for idx, k := range b[j] {
					if k == 1 {
						c[idx]++
					}
				}
			}
		}
		cnt := 0
		for j := 0; j < m; j++ {
			if c[j]%2 == p[j] {
				cnt++
			}
		}
		if cnt == m {
			ans++
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
