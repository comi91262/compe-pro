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
	var x int
	fmt.Fscan(reader, &x)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)

	ans := 0
	for i := 0; i < n; i++ {
		if x-a[i] >= 0 {
			x -= a[i]
			if i == n-1 && x > 0 {
				continue
			}
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
