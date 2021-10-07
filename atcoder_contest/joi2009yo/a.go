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
func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}

func toSecond(a []int) (x int) {
	for i := 0; i < 3; i++ {
		x *= 60
		x += a[i]
	}
	return
}

func fromSecond(x int) (a []int) {
	a = make([]int, 3)

	a[0] = x / 3600
	x %= 3600
	a[1] = x / 60
	x %= 60
	a[2] = x

	return a
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	as := scanInts(3)
	ae := scanInts(3)
	bs := scanInts(3)
	be := scanInts(3)
	cs := scanInts(3)
	ce := scanInts(3)

	a := fromSecond(toSecond(ae) - toSecond(as))
	b := fromSecond(toSecond(be) - toSecond(bs))
	c := fromSecond(toSecond(ce) - toSecond(cs))
	fmt.Fprintf(wr, "%v %v %v\n", a[0], a[1], a[2])
	fmt.Fprintf(wr, "%v %v %v\n", b[0], b[1], b[2])
	fmt.Fprintf(wr, "%v %v %v\n", c[0], c[1], c[2])

}
