package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// スライス a の i番目からj番目を反転させる関数
func reverse(a []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
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

func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	r := make([]int, 0)

	for x != 0 {
		r = append(r, mod(x, base))
		if x < 0 {
			x = (x - 1) / base
		} else {
			x = x / base
		}
	}

	reverse(r, 0, len(r)-1)

	return r
}

func mod(x, y int) int {
	x = x % y
	if x >= 0 {
		return x
	}
	if y < 0 {
		return x - y
	}
	return x + y
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for _, d := range toDigits(n, -2) {
		fmt.Fprintf(writer, "%v", d)
	}
	fmt.Fprintf(writer, "\n")

}
