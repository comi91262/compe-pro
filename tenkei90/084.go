package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type runLength struct {
	s string
	l int
}

func encodeRunLength(s []string) []runLength {
	r := make([]runLength, len(s))

	cnt := 0
	for i := 0; i < len(s); i++ {
		if cnt > 0 && r[cnt-1].s == s[i] {
			r[cnt-1].l++
		} else {
			r[cnt].s = s[i]
			r[cnt].l = 1
			cnt++
		}
	}

	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var tmp string
	fmt.Fscan(reader, &tmp)
	s := strings.Split(tmp, "")

	a := encodeRunLength(s)
	ans := 0
	for i := 0; i < len(a); i++ {
		ans += a[i].l * (n - a[i].l)
		n -= a[i].l
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
