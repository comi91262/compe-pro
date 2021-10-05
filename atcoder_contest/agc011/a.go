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
	var c int
	fmt.Fscan(reader, &c)
	var k int
	fmt.Fscan(reader, &k)

	var t = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}
	sort.Ints(t)

	ans := 1
	bus := 1
	bs := t[0]
	for i := 1; i < n; i++ {
		if t[i] <= bs+k && bus+1 <= c {
			bus++
		} else {
			bus = 1
			bs = t[i]
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
