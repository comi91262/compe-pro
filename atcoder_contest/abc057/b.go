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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
	}

	var c = make([]int, m)
	var d = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &c[i])
		fmt.Fscan(reader, &d[i])
	}

	for i := 0; i < n; i++ {
		mn := 1 << 60
		idx := 1
		for j := 0; j < m; j++ {
			dist := abs(a[i]-c[j]) + abs(b[i]-d[j])
			if dist < mn {
				mn = dist
				idx = j
			}
		}
		fmt.Fprintf(writer, "%v\n", idx+1)
	}
}


