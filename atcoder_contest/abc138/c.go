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

	var n int
	fmt.Fscan(reader, &n)
	var v = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &v[i])
	}
	sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })

	ans := 0.5 * (v[0] + v[1])
	for i := 2; i < n; i++ {
		ans = (ans + v[i]) / 2.0
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
