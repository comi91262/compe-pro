package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 998244353

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	a = a * (a + 1) / 2
	b = b * (b + 1) / 2
	c = c * (c + 1) / 2

	ans := a
	ans %= mod
	b %= mod
	ans *= b
	ans %= mod
	c %= mod
	ans *= c
	ans %= mod
	fmt.Fprintf(writer, "%v\n", ans)
}
