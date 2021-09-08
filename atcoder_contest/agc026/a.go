package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type runLength struct {
	s int
	l int
}

func encodeRunLength(s []int) []runLength {
	r := []runLength{}

	cnt := 0
	for i := 0; i < len(s); i++ {
		if cnt > 0 && r[cnt-1].s == s[i] {
			r[cnt-1].l++
		} else {
			r = append(r, runLength{s: s[i], l: 1})
			cnt++
		}
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

	e := encodeRunLength(a)
	ans := 0
	for i := range e {
		ans += e[i].l / 2
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
