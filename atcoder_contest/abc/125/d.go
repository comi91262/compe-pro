package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	p := 0
	for i := range a {
		if a[i] > 0 {
			p++
		}
	}

	ans := 0
	if (n-p)%2 == 0 {
		for i := range a {
			ans += abs(a[i])
		}
	} else {
		for i := range a {
			a[i] = abs(a[i])
		}

		sort.Ints(a)
		ans = -1 * a[0]
		for i := 1; i < n; i++ {
			ans += a[i]
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
