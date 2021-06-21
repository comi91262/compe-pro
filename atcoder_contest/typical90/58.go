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

	ans := make([]int, 0)
	for x != 0 {
		ans = append(ans, x%base)
		x = x / base
	}
	return ans
}

const d = 100000

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]int, d)
	m := n
	a[m]++
	period := 0
	for {
		period++
		sum := 0
		for _, v := range toDigits(m, 10) {
			sum += v
		}
		m = (m + sum) % d
		if a[m] > 0 {
			break
		} else {
			a[m]++
		}
	}

	var b = make([]int, d)
	b[m]++
	cycle := 0
	for {
		cycle++
		sum := 0
		for _, v := range toDigits(m, 10) {
			sum += v
		}
		m = (m + sum) % d
		if b[m] > 0 {
			break
		} else {
			b[m]++
		}
	}

	if k > period {
		k = (k-period)%cycle + period
	}

	for k > 0 {
		sum := 0
		for _, v := range toDigits(n, 10) {
			sum += v
		}
		n = (n + sum) % d
		k--
	}
	fmt.Fprintf(writer, "%d\n", n)
}
