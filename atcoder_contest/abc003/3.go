package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var r = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &r[i])
	}

	sort.Slice(r, func(i, j int) bool { return r[i] > r[j] })

	s := make([]int, k)
	copy(s, r)
	// fmt.Fprintf(writer, "%v\n", s)
	sort.Ints(s)

	ans := 0.0

	for i := 0; i < k; i++ {
		ans = (ans + float64(s[i])) / 2.0
		// fmt.Fprintf(writer, "%v\n", ans)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}

