package main

import (
	"bufio"
	"fmt"
	"os"
)

const w = 1001

var a [w + 1][w + 1]int

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var lx, ly, rx, ry int
		fmt.Fscan(reader, &lx, &ly, &rx, &ry)

		a[lx][ly]++
		a[lx][ry]--
		a[rx][ly]--
		a[rx][ry]++
	}

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			a[i][j+1] += a[i][j]
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			a[i+1][j] += a[i][j]
		}
	}

	b := make([]int, n+1)
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] > 0 {
				b[a[i][j]]++
			}
		}
	}

	for i := 1; i < n+1; i++ {
		fmt.Fprintf(writer, "%d\n", b[i])
	}

}

