package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int       { sc.Scan(); return parseInt(sc.Bytes()) }
func scanString() string { sc.Scan(); return string(sc.Bytes()) }
func scanFloat() float64 { sc.Scan(); return parseFloat(sc.Bytes()) }
func scanInts(n int) (ints []int) {
	ints = make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = scanInt()
	}
	return
}
func scanStrings(n int) (st []string) {
	st = make([]string, n)
	for i := 0; i < n; i++ {
		st[i] = scanString()
	}
	return
}
func scanFloats(n int) (fs []float64) {
	fs = make([]float64, n)
	for i := 0; i < n; i++ {
		fs[i] = scanFloat()
	}
	return fs
}

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

var float64pow10 = []float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6,
	1e7, 1e8, 1e9, 1e10, 1e11, 1e12,
	1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
	1e19, 1e20, 1e21, 1e22,
}

func parseFloat(b []byte) float64 {
	f, dot := 0.0, 0
	for i, ch := range b {
		if ch == '.' {
			dot = i + 1
			continue
		}
		if ch -= '0'; ch <= 9 {
			f = f*10 + float64(ch)
		}
	}
	if dot != 0 {
		f /= float64pow10[len(b)-dot]
	}
	if b[0] == '-' {
		return -f
	}
	return f
}

type pair struct {
	first  int
	second int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()
	p := make([]pair, m)
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = scanInt()
		b[i] = scanInt()
		p[i].first = a[i]
		p[i].second = b[i]
	}
	sort.Slice(p, func(i, j int) bool { return p[i].first > p[j].first })
	ans := 0
	for i := 0; i < m-1; i++ {
		if p[i].first < n {
			ans += n - p[i].first
		}
	}
	fmt.Fprintf(wr, "%v\n", ans)
}
