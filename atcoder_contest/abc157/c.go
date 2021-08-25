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
	var m int
	fmt.Fscan(reader, &m)

	var s = make([]int, m)
	var c = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &c[i])
	}
	l := [3]int{0, 10, 100}
	r := [3]int{9, 99, 999}

	for i := l[n-1]; i <= r[n-1]; i++ {
		flag := true

		for j := 0; j < m; j++ {
			if i/pow(10, n-s[j])%10 != c[j] {
				flag = false
				break
			}
		}
		if flag {
			fmt.Fprintf(writer, "%v\n", i)
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", -1)
	return
}
