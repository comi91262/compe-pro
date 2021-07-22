package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

// 前処理　O(n)
// クエリ  O(1)
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

var nck = FastNck{}

func calc(x, y, dl int) int {
	if x <= 0 || y <= 0 {
		return 0
	}

	ans := nck.Combination(x*y, dl)
	ans %= mod

	return ans
}

func main() {
	defer writer.Flush()

	var r, c int
	fmt.Fscan(reader, &r, &c)
	var x, y int
	fmt.Fscan(reader, &x, &y)
	var d, l int
	fmt.Fscan(reader, &d, &l)

	nck.Create(31*31, mod)

	ans := calc(x, y, d+l)
	ans %= mod
	ans -= (2*calc(x-1, y, d+l) + 2*calc(x, y-1, d+l))
	ans %= mod
	ans += calc(x, y-2, d+l) + calc(x-2, y, d+l) + 4*calc(x-1, y-1, d+l)
	ans %= mod
	ans -= (2*calc(x-1, y-2, d+l) + 2*calc(x-2, y-1, d+l))
	ans %= mod
	ans += calc(x-2, y-2, d+l)
	ans %= mod

	ans = (ans + mod) % mod
	ans *= calc(d+l, 1, l)
	ans %= mod

	ans *= (r - x + 1)
	ans %= mod
	ans *= (c - y + 1)
	ans %= mod

	fmt.Fprintf(writer, "%v\n", ans)

}
