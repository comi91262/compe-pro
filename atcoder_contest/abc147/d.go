package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	for i := 0; i < 60; i++ {
		b := 1 << i

		one, zero := 0, 0
		for j := 0; j < n; j++ {
			if b&a[j] != 0 {
				one++
			} else {
				zero++
			}
		}

		ans += (((b % mod) * one % mod) * zero) % mod
		ans %= mod
	}

	fmt.Fprintf(writer, "%v\n", ans%mod)
}
