package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

// func fact(n int) int {
// 	if n == 1 {
// 		return 1
// 	} else {
// 		return n * fact(n-1)
// 	}
// }
//
// func combination(n int, r int) int {
// 	return fact(n) / fact(r) * fact(n-r)
// }

// var b [1_000_000_000]int

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := combination(n, 2)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	sum := 0
	for _, v := range m {
		sum += combination(v, 2)

	}
	//	for i := 1; i <= len(m); i++ {
	//		if m[i] == 1 || m[i] == 0 {
	//			continue
	//		}
	//	}

	fmt.Fprintf(writer, "%d\n", ans-sum)
}
