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

func intersection(a, b int) int {
	return a & b
}
func union(a, b int) int {
	return a | b
}
func diff(a, b int) int {
	return a ^ b
}
func subtract(a, x int) int {
	b := 1 << (49 - x)
	return a &^ b
}
func increment(a int) int {
	if a&1<<0 != 0 {
		return a>>1 | 1<<49
	} else {
		return a >> 1
	}
}
func decrement(a int) int {
	if a&1<<49 != 0 {
		return a<<1&(1<<50-1) | 1<<0
	} else {
		return a << 1 & (1<<50 - 1)
	}
}

func print(s int) {
	out := fmt.Sprintf("%050b", s)
	outs := []int{}
	for i := 0; i < len(out); i++ {
		if out[i] == '1' {
			outs = append(outs, i)
		}
	}

	for i := range outs {
		if i > 0 {
			fmt.Fprintf(wr, " ")
		}
		fmt.Fprintf(wr, "%v", outs[i])
	}
	fmt.Fprintf(wr, "\n")
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	a, b := 0, 0
	n := scanInt()
	for i := 0; i < n; i++ {
		x := scanInt()
		a |= 1 << (49 - x)
	}

	m := scanInt()
	for i := 0; i < m; i++ {
		x := scanInt()
		b |= 1 << (49 - x)
	}

	switch scanString() {
	case "intersection":
		print(intersection(a, b))
	case "union_set":
		print(union(a, b))
	case "symmetric_diff":
		print(diff(a, b))
	case "subtract":
		x := scanInt()
		print(subtract(a, x))
	case "increment":
		print(increment(a))
	case "decrement":
		print(decrement(a))
	}
}
