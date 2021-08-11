package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	ans := 0
	for i := range a {
		if i%2 == 0 {
			ans += a[i] * a[i]
		} else {
			ans -= a[i] * a[i]
		}
	}
	fmt.Fprintf(writer, "%v\n", math.Pi*float64(ans))
}
