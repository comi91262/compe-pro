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

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()

	var w int
	fmt.Fscan(reader, &w)
	var h int
	fmt.Fscan(reader, &h)

	nck := FastNck{}
	nck.Create(w+h-2, mod)

	fmt.Fprintf(writer, "%v\n", nck.Combination(w+h-2, w-1))

}
