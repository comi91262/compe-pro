package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	mx := 0
	var b = []int{}
	for k, v := range m {
		if v >= 4 {
			mx = max(mx, k*k)
			b = append(b, k)
		}
		if v >= 2 {
			b = append(b, k)
			continue
		}
	}
	if len(b) < 1 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })

	fmt.Fprintf(writer, "%v\n", max(mx, b[0]*b[1]))
}
