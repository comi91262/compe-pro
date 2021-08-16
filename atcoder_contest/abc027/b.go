package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	total := sum(a)
	if total%n != 0 {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}

	ans := 0
	m := total / n
	for i := 0; i < n; i++ {
		cnt := 1
		tmp := a[i]
		for tmp != m*cnt {
			ans++

			tmp += a[i+cnt]
			cnt++
		}
		for j := i; j < i+cnt; j++ {
			a[j] = m
		}

		i += cnt - 1
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
