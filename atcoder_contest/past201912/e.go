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
	var q int
	fmt.Fscan(reader, &q)

	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	for i := 0; i < q; i++ {
		var t int
		var a, b int
		fmt.Fscan(reader, &t)
		switch t {
		case 1:
			fmt.Fscan(reader, &a, &b)
			a--
			b--
			f[a][b] = 1
		case 2:
			fmt.Fscan(reader, &a)
			a--
			for j := 0; j < n; j++ {
				if f[j][a] == 1 {
					f[a][j] = 1
				}
			}
		case 3:
			fmt.Fscan(reader, &a)
			a--
			cand := []int{}
			for j := range f[a] {
				if f[a][j] == 1 {
					cand = append(cand, j)
				}
			}

			for _, j := range cand {
				for k := range f[j] {
					if f[j][k] == 1 && a != k {
						f[a][k] = 1
					}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if f[i][j] == 1 {
				fmt.Fprintf(writer, "Y")
			} else {
				fmt.Fprintf(writer, "N")
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}
