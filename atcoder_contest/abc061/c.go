package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type data struct {
	n   int
	cnt int
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	var a = make([]int, n)
	var b = make([]int, n)
	var c = make([]data, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		c[i].n = a[i]
		c[i].cnt = b[i]
	}

	sort.Slice(c, func(i, j int) bool { return c[i].n < c[j].n })

	cnt := 0
	for i := 0; i < n; i++ {
		if cnt <= k && k <= cnt+c[i].cnt {
			fmt.Fprintf(writer, "%v\n", c[i].n)
			break
		}
		cnt += c[i].cnt
	}

}
