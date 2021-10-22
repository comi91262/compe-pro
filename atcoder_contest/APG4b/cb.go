package main

import (
	"bufio"
	"fmt"
	"os"
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

func scanPairInts(n int) (a, b []int) {
	a = make([]int, n)
	b = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		b[i] = scanInt()
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

func f0(n int) int {
	return 1
}
func f1(n, m int) int {
	s := 0
	for i := 0; i < n; i++ {
		s++
	}
	for i := 0; i < m; i++ {
		s++
	}
	return s
}
func f2(n int) int {
	s := 0
	for i := 0; i < n; i++ {
		t := n
		cnt := 0
		for t > 0 {
			cnt++
			t /= 2
		}
		s += cnt
	}
	return s
}

func f3(n int) int {
	s := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s++
		}
	}

	return s
}

func f4(n int) int {
	return -1
}

func f5(n, m int32) int32 {
	var s int32
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < m; j++ {
			s += i + j
		}
	}
	return s
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	n := scanInt()
	m := scanInt()

	fmt.Fprintf(wr, "f0: %v\n", f0(n))
	fmt.Fprintf(wr, "f1: %v\n", f1(n, m))
	fmt.Fprintf(wr, "f2: %v\n", f2(n))
	fmt.Fprintf(wr, "f3: %v\n", f3(n))
	fmt.Fprintf(wr, "f4: %v\n", f4(n))
	fmt.Fprintf(wr, "f5: %v\n", f5(int32(n), int32(m)))
}
