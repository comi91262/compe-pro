package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var a [101][101]int
var b [101][101]int

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()

	var h, w int
	fmt.Fscan(reader, &h, &w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &b[i][j])
		}
	}

	cnt := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			v := b[i][j] - a[i][j]
			cnt += abs(v)
			a[i][j] += v
			a[i+1][j] += v
			a[i][j+1] += v
			a[i+1][j+1] += v
		}
	}

	flag := true

	for i := 0; i < h; i++ {
		flag = flag && a[i][w-1] == b[i][w-1]
	}

	for i := 0; i < w; i++ {
		flag = flag && a[h-1][i] == b[h-1][i]
	}

	if flag {
		fmt.Fprintf(writer, "Yes\n")
		fmt.Fprintf(writer, "%v\n", cnt)
	} else {
		fmt.Fprintf(writer, "No\n")
	}
}
