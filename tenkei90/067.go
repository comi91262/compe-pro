package main

import (
	"bufio"
	"fmt"
	"os"
)

// 123 -> [3,2,1] (10)
func strToDigits(s string) []int {
	ans := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		var n = int(s[i] - 48)
		ans[len(s)-1-i] = n
	}
	return ans
}

// 123 -> [3,2,1] (10)
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

func toNumber(a []int, base int) int {
	cnt := 1
	ans := 0
	for i := 0; i < len(a); i++ {
		ans += a[i] * cnt
		cnt *= base
	}
	return ans
}

func convert(a []int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] == 8 {
			a[i] = 5
		}
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n string
	var k int
	fmt.Fscan(reader, &n, &k)

	x := strToDigits(n)
	for i := 0; i < k; i++ {
		x = convert(toDigits(toNumber(x, 8), 9))
	}
	ans := toNumber(x, 10)
	fmt.Fprintf(writer, "%d\n", ans)
}
