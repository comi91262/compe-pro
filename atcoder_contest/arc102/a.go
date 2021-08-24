package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	m := map[int]int{}
	for i := 1; i <= n; i++ {
		m[i%k]++
	}
	ans := 0
	// fmt.Fprintf(writer, "%v\n", m)
	for i, v := range m {
		if (2*i)%k != 0 {
			continue
		}
		ans += pow(v, 3)
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
