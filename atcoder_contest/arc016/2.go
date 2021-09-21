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

func encodeRunLength(s []string) (r []runLength) {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if 0 < cnt && r[cnt-1].s == s[i] && s[i] != "x" {
			r[cnt-1].l++
		} else {
			r = append(r, runLength{s: s[i], l: 1})
			cnt++
		}
	}
	return
}
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	s := make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	t := make([][]string, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < n; j++ {
			t[i] = append(t[i], s[j][i])
		}
	}

	ans := 0
	for i := 0; i < 9; i++ {
		for _, v := range encodeRunLength(t[i]) {
			if v.s == "o" || v.s == "x" {
				ans++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
