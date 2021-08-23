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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	e := make([]int, n)
	w := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == "E" {
			e[i]++
		} else {
			w[i]++
		}
	}

	for i := 1; i < n; i++ {
		e[i] += e[i-1]
		w[i] += w[i-1]
	}

	ans := 1 << 60
	for i := 0; i < n; i++ {
		total := 0
		if i >= 1 {
			total += w[i-1]
		}
		total += e[n-1] - e[i]

		ans = min(ans, total)
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
