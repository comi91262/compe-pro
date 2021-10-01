package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	r := make([]int, 0)
	for x != 0 {
		r = append(r, x%base)
		x = x / base
	}

	return r
}
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

	ans := 0
	for {
		a := toDigits(n, 10)

		if len(a) < 4 {
			break
		}

		comma := (len(a)+2)/3 - 1
		diff := n - (pow(10, len(a)-1) - 1)
		ans += comma * diff
		n -= diff
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	ans := 0
	if n >= 1000 {
		ans += n - 1000 + 1
	}
	if n >= 1000000 {
		ans += n - 1000000 + 1
	}
	if n >= 1000000000 {
		ans += n - 1000000000 + 1
	}
	if n >= 1000000000000 {
		ans += n - 1000000000000 + 1
	}
	if n >= 1000000000000000 {
		ans += n - 1000000000000000 + 1
	}

	fmt.Fprintf(writer, "%v\n", ans)
}

