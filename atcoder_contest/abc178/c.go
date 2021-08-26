package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type FastNck struct {
	fact    []int
	factInv []int
	prime   int
}

func (table *FastNck) Create(n, p int) {
	table.fact = make([]int, n+1)
	table.factInv = make([]int, n+1)
	table.prime = p

	table.fact[0], table.fact[1] = 1, 1
	table.factInv[0], table.factInv[1] = 1, 1

	inv := make([]int, n+1)
	inv[1] = 1
	for i := 2; i < n+1; i++ {
		table.fact[i] = table.fact[i-1] * i % p
		inv[i] = p - inv[p%i]*(p/i)%p
		table.factInv[i] = table.factInv[i-1] * inv[i] % p
	}
}

func (table *FastNck) Combination(n, k int) int {
	if n < k {
		return 0
	}
	if k == 0 {
		return 1
	}

	return table.fact[n] * (table.factInv[k] * table.factInv[n-k] % table.prime) % table.prime
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

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	c := FastNck{}
	c.Create(n, mod)

	ans := 0
	for i := 2; i <= n; i++ {
		inter := c.Combination(n, i)
		inter %= mod

		inter *= (powMod(2, i, mod) - 2 + mod)
		inter %= mod

		inter *= powMod(8, n-i, mod)
		inter %= mod

		ans += inter
		ans %= mod
	}
	fmt.Fprintf(writer, "%v\n", ans%mod)
}
