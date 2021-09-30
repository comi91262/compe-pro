package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	m := map[int]int{}
	b := make([]map[int]int, n)

	pf := PrimeFactor{}
	pf.Constitute(1_000_000)

	for i := 0; i < n; i++ {
		b[i] = pf.Do(a[i])
	}

	for i := 0; i < n; i++ {
		for p, cnt := range b[i] {
			m[p] += cnt
		}
	}

	pc := true
	for i := 0; i < n; i++ {
		for p, cnt := range b[i] {
			if m[p]-2*cnt >= 0 {
				pc = false
				break
			}
		}
	}
	if pc {
		fmt.Fprintf(writer, "%v\n", "pairwise coprime")
		return
	}

	g := 0
	for i := 0; i < n; i++ {
		g = gcd(g, a[i])
	}
	if g == 1 {
		fmt.Fprintf(writer, "%v\n", "setwise coprime")
		return
	}

	fmt.Fprintf(writer, "%v\n", "not coprime")
}
