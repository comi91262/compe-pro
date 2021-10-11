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

func makeNumberArray(size, depth, lower, upper int, memo map[int]int, a []int, accum *[][]int) {
	if size == depth {
		b := make([]int, len(a))
		copy(b, a)
		*accum = append(*accum, b)
		return
	}

	for i := lower; i <= upper; i++ {
		if memo[i] > 0 {
			continue
		}
		a = append(a, i)
		memo[i]++
		makeNumberArray(size, depth+1, lower, upper, memo, a, accum)
		a = a[:len(a)-1]
		memo[i]--
	}
}
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	n := scanInt()
	m := scanInt()
	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = scanInts(m)
	}
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = scanInts(n)
		for j := 0; j < n; j++ {
			c[i][j]--
		}
	}
	s := map[int]int{}
	t := map[int]int{}
	u := map[int]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch (i + j) % 3 {
			case 0:
				s[c[i][j]]++
			case 1:
				t[c[i][j]]++
			case 2:
				u[c[i][j]]++
			}
		}
	}

	//a := make([]map[int]int, m)
	//for i := 0; i < m; i++ {
	//	a[i] = map[int]int{}
	//	for j := 0; j < m; j++ {
	//		a[i][d[j][i]]++
	//	}
	//}
	b := [][]int{}
	makeNumberArray(3, 0, 0, m-1, map[int]int{}, []int{}, &b)
	ans := 1 << 60
	for _, arr := range b {
		i, j, k := arr[0], arr[1], arr[2]
		sum := 0
		for b, cnt := range s {
			sum += d[b][i] * cnt
		}
		for b, cnt := range t {
			sum += d[b][j] * cnt
		}
		for b, cnt := range u {
			sum += d[b][k] * cnt
		}
		ans = min(ans, sum)

	}

	fmt.Fprintf(wr, "%v\n", ans)
}
