package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type FastNck2 struct {
	factInv []int
	prime   int
}

func (table *FastNck2) Create(k, p int) {
	table.factInv = make([]int, k+1)
	table.prime = p
	inv := make([]int, k+1)

	table.factInv[0], table.factInv[1] = 1, 1
	inv[1] = 1

	for i := 2; i < k; i++ {
		inv[i] = p - inv[p%i]*(p/i)%p
		table.factInv[i] = table.factInv[i-1] * inv[i] % p
	}
}

func (table *FastNck2) Combination(n, k int) int {
	r := 1
	for i := n; i >= n-k+1; i-- {
		r *= i
		r %= table.prime
	}
	return r * table.factInv[k] % table.prime
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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	nck := FastNck2{}
	nck.Create(200010, mod)

	an := nck.Combination(n, a)
	bn := nck.Combination(n, b)

	ans := (powMod(2, n, mod) - 1 + mod) % mod
	ans = (ans - an + mod) % mod
	ans = (ans - bn + mod) % mod

	if ans < 0 {
		fmt.Fprintf(writer, "%v\n", 0)
	} else {
		fmt.Fprintf(writer, "%v\n", ans)
	}
}
