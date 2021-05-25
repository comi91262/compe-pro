package main

import (
	"bufio"
	"fmt"
	"os"
)

var a [100001]struct{ c, p int }
var b [100001]struct{ l, r int }
var s1 [100001]int
var s2 [100001]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i].c)
		fmt.Fscan(reader, &a[i].p)
	}

	var sum1, sum2 int
	for i := 0; i <= n; i++ {
		if a[i].c == 1 {
			sum1 += a[i].p
		} else {
			sum2 += a[i].p
		}
		s1[i] = sum1
		s2[i] = sum2
	}

	var q int
	fmt.Fscan(reader, &q)
	for i := 1; i <= q; i++ {
		fmt.Fscan(reader, &b[i].l)
		fmt.Fscan(reader, &b[i].r)
	}

	for i := 1; i <= q; i++ {
		fmt.Fprintf(writer, "%d %d\n", s1[b[i].r]-s1[b[i].l-1], s2[b[i].r]-s2[b[i].l-1])
	}
}
