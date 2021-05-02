package main

import (
	"fmt"
	"math"
)

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count += 1
	}
	return count
}

func intToDigits(n int) []int {
	c := countDigits(n)
	result := make([]int, c)
	for i := 0; i < c; i++ {
		r := n % int(math.Pow10(c-i))
		result[i] = r / int(math.Pow10(c-i-1))
	}

	return result
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	count := 0
	for i := 1; i < n+1; i++ {
		digits := intToDigits(i)
		// fmt.Printf("%v", digits)
		sum := 0
		for _, d := range digits {
			sum += d
		}
		// println(sum)
		if a <= sum && sum <= b {
			count += i
		}
	}

	fmt.Println(count)
}
