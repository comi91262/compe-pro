package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_000 + 7

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var t = make([]int, n)
	var l = make([]float64, n)
	var r = make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i], &l[i], &r[i])
	}

	for i := 0; i < n; i++ {
		switch t[i] {
		case 1:
		case 2:
			r[i] -= 0.01
		case 3:
			l[i] += 0.01
		case 4:
			r[i] -= 0.01
			l[i] += 0.01
		}

	}

	sum := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if r[i] >= l[j] && r[j] >= l[i] {
				sum++
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", sum)
}
