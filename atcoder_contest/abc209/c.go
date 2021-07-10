package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	sort.Ints(c)

	ans := 1
	cnt := 0
	for i := 0; i < n; i++ {
		ans *= c[i] - cnt
		ans %= mod
		cnt++

		if ans == 0 {
			break
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
