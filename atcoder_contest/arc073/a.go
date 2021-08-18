package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var t int
	fmt.Fscan(reader, &t)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			ans += t
			continue
		}

		ans += t
		if a[i-1]+t > a[i] {
			ans -= a[i-1] + t - a[i]
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
