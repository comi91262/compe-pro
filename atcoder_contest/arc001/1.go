package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
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
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	m := map[string]int{}

	for _, v := range s {
		m[v]++
	}

	mn := 1 << 60
	mx := 0
	for _, v := range m {
		mn = min(mn, v)
		mx = max(mx, v)
	}
	if len(m) != 4 {
		mn = 0
	}
	fmt.Fprintf(writer, "%v %v\n", mx, mn)
}
