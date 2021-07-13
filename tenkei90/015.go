package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

// nCk = n! / k! / (n-k)! (mod p)
//     = n! * (k!)^(-1) * ((n-k)!)^(-1) (mod p)
//     = fact[n] * factInv[k] * factInv[n-k] (mod p)
//
// fact[i] = i * fact[i-1] (mod p)
// factInv[i] = i!^(-1) = inv[i] * fact_inv[i-1] (mod p, inv[i] = (i)^(-1))
func initFactTb(n, p int) ([]int, []int) {
	fact := make([]int, n+1)
	factInv := make([]int, n+1)

	fact[0], fact[1], factInv[0], factInv[1] = 1, 1, 1, 1

	inv := make([]int, n+1)
	inv[1] = 1
	for i := 2; i < n+1; i++ {
		fact[i] = fact[i-1] * i % p
		inv[i] = p - inv[p%i]*(p/i)%p
		factInv[i] = factInv[i-1] * inv[i] % p
	}
	return fact, factInv
}

func combinationMod(n, k, p int, fact, factInv []int) int {
	if n < k {
		return 0
	}
	if k == 0 {
		return 1
	}

	return fact[n] * (factInv[k] * factInv[n-k] % p) % p
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	fact, factInv := initFactTb(n, mod)

	for k := 1; k <= n; k++ {
		ans := 0
		for a := 1; a <= (n/k + 1); a++ {
			ans += combinationMod(n-(k-1)*(a-1), a, mod, fact, factInv)
			ans %= mod
		}
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
