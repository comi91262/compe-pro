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
func powMod(a, x, d int) int {
	var r int = 1
	for x > 0 {
		if x&1 == 1 {
			r = r * a % d
		}
		a = a * a % d
		x >>= 1
	}
	return r
}

type PrimeFactor struct {
	spf []int
}

func (p *PrimeFactor) Constitute(n int) {
	p.spf = make([]int, n+1)

	for i := 0; i <= n; i++ {
		p.spf[i] = i
	}
	for i := 2; i*i <= n; i++ {
		if p.spf[i] == i {
			for j := i * i; j <= n; j += i {
				if p.spf[j] == j {
					p.spf[j] = i
				}
			}
		}
	}
}

func (p *PrimeFactor) Do(n int) map[int]int {
	var m = map[int]int{}
	for n != 1 {
		m[p.spf[n]]++
		n /= p.spf[n]
	}
	return m
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

const mod = 1_000_000_000 + 7

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	n := scanInt()
	a := scanInts(n)

	p := PrimeFactor{}
	p.Constitute(1000000)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		for p, cnt := range p.Do(a[i]) {
			m[p] = max(m[p], cnt)
		}
	}

	ans := 0
	l := 1
	for p, cnt := range m {
		l *= powMod(p, cnt, mod)
		l %= mod
	}
	for i := 0; i < n; i++ {
		ans += l * powMod(a[i], mod-2, mod) % mod
		ans %= mod
	}
	fmt.Fprintf(wr, "%v\n", ans)
}
