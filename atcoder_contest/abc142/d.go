package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func primeFactor(n int) map[int]int {
	var m = map[int]int{}

	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n]++
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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	g := gcd(a, b)
	p := primeFactor(g)

	fmt.Fprintf(writer, "%v\n", len(p)+1)
}
