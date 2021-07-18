package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func toBinaryDigits(x, size int) []int {
	r := make([]int, size)
	for i := 0; i < size; i++ {
		if x == 0 {
			break
		}
		r[size-1-i] = x % 2
		x /= 2
	}
	return r
}

func main() {
	defer writer.Flush()
	var n, q int
	fmt.Fscan(reader, &n, &q)

	var x = make([]int, q)
	var y = make([]int, q)
	var z = make([]int, q)
	var w = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &x[i], &y[i], &z[i], &w[i])
	}

	ans := 1
	for bit := 0; bit < 60; bit++ {
		ways := 0
		for i := 0; i < 1<<n; i++ {
			digits := toBinaryDigits(i, n)
			cnt := 0
			for j := 0; j < q; j++ {
				wd := toBinaryDigits(w[j], 60)
				if wd[bit] == digits[x[j]-1]|digits[y[j]-1]|digits[z[j]-1] {
					cnt++
				}
			}
			if cnt == q {
				ways++
			}
		}
		ans *= ways
		ans %= mod
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
