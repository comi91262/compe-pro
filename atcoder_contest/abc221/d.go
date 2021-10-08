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
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
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
func unique(a []int) []int {
	r := make([]int, 0)
	m := make(map[int]bool, 0)

	for _, e := range a {
		if !m[e] {
			m[e] = true
			r = append(r, e)
		}
	}
	return r
}

const mod = 1_000_000_000 + 7

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	n := scanInt()

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		b[i] = a[i] + scanInt()
	}

	e := []int{}
	for i := 0; i < n; i++ {
		e = append(e, a[i])
		e = append(e, b[i])
	}
	sort.Ints(e)
	e = unique(e)
	packed := map[int]int{}
	unpacked := map[int]int{}
	for i := 0; i < n; i++ {
		packed[a[i]] = sort.Search(len(e), func(j int) bool { return a[i] <= e[j] })
		packed[b[i]] = sort.Search(len(e), func(j int) bool { return b[i] <= e[j] })
		unpacked[packed[a[i]]] = a[i]
		unpacked[packed[b[i]]] = b[i]
	}

	f := make([]int, 2*n+1)
	for i := 0; i < n; i++ {
		f[packed[a[i]]] += 1
		f[packed[b[i]]] -= 1
	}
	for i := 0; i < 2*n; i++ {
		f[i+1] += f[i]
	}

	m := map[int]int{}
	for i := 0; i < 2*n; i++ {
		m[f[i]] += unpacked[i+1] - unpacked[i]
	}
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Fprintf(wr, " ")
		}
		fmt.Fprintf(wr, "%v", m[i])
	}
	fmt.Fprintf(wr, "\n")

}


