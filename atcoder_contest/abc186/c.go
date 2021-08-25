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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	ans := 0
	for i := 1; i <= n; i++ {
		flag := true
		for _, d := range toDigits(i, 10) {
			if d == 7 {
				flag = false
			}
		}
		for _, d := range toDigits(i, 8) {
			if d == 7 {
				flag = false
			}
		}
		if flag {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
