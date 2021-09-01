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
	m := a[0]
	b := a[1:]

	ans := 0
	mn := float64(1 << 60)
	for i := 0; i < len(b); i++ {
		d := math.Abs(float64(b[i]) - float64(m)/2.0)
		if mn > d {
			mn = d
			ans = i
		}
	}
	fmt.Fprintf(writer, "%v %v\n", m, b[ans])
}
