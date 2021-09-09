package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func combination(n, k int) int {
	r := 1
	for d := 1; d <= k; d++ {
		r *= n
		n--
		r /= d
	}

	return r
}

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var k int
	fmt.Fscan(reader, &k)

	ans := ""
	ab := a + b
	total := 0

	for ab > 0 {
		c := combination(ab-1, b)
		if total+c < k {
			ans += "b"
			b--
			total += c
		} else {
			ans += "a"
			a--
		}
		ab--
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
