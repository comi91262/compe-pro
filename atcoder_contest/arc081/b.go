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

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var _ = strings.Split(ss, "")

	e := encodeRunLength(s)

	ans := 0
	if e[0].l%2 == 0 {
		ans = 6 % mod
	} else {
		ans = 3 % mod
	}
	for i := 0; i < len(e)-1; i++ {
		switch {
		case e[i].l%2 == 0 && e[i+1].l%2 == 0:
			ans *= (3 % mod)
		case e[i].l%2 == 1 && e[i+1].l%2 == 0:
			ans *= (2 % mod)
		case e[i].l%2 == 0 && e[i+1].l%2 == 1:
			ans *= (1 % mod)
		case e[i].l%2 == 1 && e[i+1].l%2 == 1:
			ans *= (2 % mod)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans%mod)
}
