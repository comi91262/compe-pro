package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	b := strings.Split(strings.Join(a, ""), "")
	c := make([]int, len(b))
	for i := 0; i < len(b); i++ {
		c[i], _ = strconv.Atoi(b[i])
	}

	ans := 0
	for i := 0; i < len(c); i++ {
		ans += c[len(c)-i-1] * powMod(10, i, mod)
		ans %= mod
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
