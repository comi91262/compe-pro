package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int       { sc.Scan(); return parseInt(sc.Bytes()) }
func scanString() string { sc.Scan(); return string(sc.Bytes()) }

func parseInt(b []byte) (n int) {
	for _, ch := range b {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
		}
	}
	if b[0] == '-' {
		return -n
	}
	return
}

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
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

type queue []string

func (q *queue) push(n string) {
	*q = append(*q, n)
}

func (q *queue) pop() string {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) front() string {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
}
func chmax(x *int, y int) {
	*x = max(*x, y)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)
	s := strings.Split(scanString(), "")
	k := scanInt()
	q := queue{}

	ans := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		q.push(s[i])
		switch s[i] {
		case ".":
			cnt++
		}
		for !q.empty() && k < cnt {
			t := q.pop()
			switch t {
			case ".":
				cnt--
			}
		}

		chmax(&ans, len(q))
	}
	fmt.Fprintf(wr, "%v\n", ans)
}
