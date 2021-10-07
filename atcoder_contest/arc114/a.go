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
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
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
func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}

func makePrimes(n int) (primes []int) {
	prime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		prime[i] = true
	}

	prime[0], prime[1] = false, false
	for i := 2; i*i <= n; i++ {
		if !prime[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			prime[j] = false
		}
	}

	for i := 2; i <= n; i++ {
		if !prime[i] {
			continue
		}

		primes = append(primes, i)
	}
	return
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n := scanInt()
	a := scanInts(n)

	p := makePrimes(50)

	ans := 1 << 60
	for i := 0; i < 1<<len(p); i++ {
		y := 1
		for j := 0; j < len(p); j++ {
			if i&(1<<j) != 0 {
				y *= p[j]
			}
		}
		if y == 1 {
			continue
		}
		isOk := true
		for j := 0; j < n; j++ {
			if gcd(a[j], y) == 1 {
				isOk = false
			}
		}
		if isOk {
			ans = min(ans, y)
		}
	}

	fmt.Fprintf(wr, "%v\n", ans)
}