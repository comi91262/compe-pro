package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

	var s string
	fmt.Fscan(reader, &s)
	n := len(s)

	ans, tot := 0, 0
	m := map[int]int{0: 1}
	for i := 0; i < n; i++ {
		tot += int(s[n-i-1]-"0"[0]) * powMod(10, i, 2019)
		tot %= 2019
		ans += m[tot]
		m[tot]++
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
