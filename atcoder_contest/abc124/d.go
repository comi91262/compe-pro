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

type queue []runLength

func (q *queue) push(n runLength) {
	*q = append(*q, n)
}

func (q *queue) pop() runLength {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	e := encodeRunLength(s)
	q := queue{}

	cnt := 0
	l := 0
	ans := 0
	for i := 0; i < len(e); i++ {
		q.push(e[i])
		l += e[i].l
		if e[i].s == "0" {
			cnt++
		}

		for !q.empty() && cnt > k {
			ee := q.pop()
			l -= ee.l
			if ee.s == "0" {
				cnt--
			}
		}

		ans = max(ans, l)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
