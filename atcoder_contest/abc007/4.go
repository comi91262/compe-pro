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

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// スライス a の i番目からj番目を反転させる関数
func reverse(a []int, i, j int) {
	for k := 0; k < (j-i+1)/2; k++ {
		swap(a, i+k, j-k)
	}
}

func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	r := make([]int, 0)
	for x != 0 {
		r = append(r, x%base)
		x = x / base
	}
	reverse(r, 0, len(r)-1)

	return r
}

func toNumber(a []int, base int) int {
	cnt := 1
	ans := 0
	for i := 0; i < len(a); i++ {
		ans += a[len(a)-1-i] * cnt
		cnt *= base
	}
	return ans
}

func calc(x int, dp [][]int) int {
	digits := toDigits(x, 10)
	n := len(digits)
	r := 0

	// fmt.Fprintf(writer, "%v\n", digits)
	for i := range digits {
		// fmt.Fprintf(writer, "%v\n", r)
		if digits[i] == 0 {
			continue
		}
		r += dp[n-i][digits[i]-1]

		if digits[i] == 4 || digits[i] == 9 {
			r += 1 + toNumber(digits[i+1:n], 10)
			break
		}

	}
	// fmt.Fprintf(writer, "%v\n", r)

	return r
}

func isBannedNumber(x int) bool {
	digits := toDigits(x, 10)
	for i := range digits {
		if digits[i] == 4 || digits[i] == 9 {
			return true
		}
	}

	return false
}

func main() {
	defer writer.Flush()

	var a, b int
	fmt.Fscan(reader, &a, &b)

	if a == 1000000000000000000 {
		a = 1000000000000000000 - 1
	}
	if b == 1000000000000000000 {
		b = 1000000000000000000 - 1
	}

	dp := make([][]int, 19)
	for i := 0; i < 19; i++ {
		dp[i] = make([]int, 10)
	}

	dp[1][4] = 1
	dp[1][9] = 1
	for i := 2; i < 19; i++ {
		for j := 0; j < 10; j++ {
			if j == 4 || j == 9 {
				dp[i][j] = pow(10, i-1)
			} else {
				for k := 0; k < 10; k++ {
					dp[i][j] += dp[i-1][k]
				}
			}
		}
	}
	for i := 1; i < 19; i++ {
		for j := 1; j < 10; j++ {
			dp[i][j] += dp[i][j-1]
		}
	}

	// for i := 1; i < 19; i++ {
	// 	fmt.Fprintf(writer, "%v\n", dp[i])
	// }

	av := calc(a-1, dp)
	bv := calc(b, dp)

	fmt.Fprintf(writer, "%v\n", bv-av)
}
