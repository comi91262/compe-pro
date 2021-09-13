package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	if n == 1 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}
	fmt.Fprintf(writer, "%v\n", len(makePrimes(n-1)))
}
