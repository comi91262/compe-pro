package main

import (
	"bufio"
	"fmt"
	"os"
)

const mult = 46

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	var b = make([]int, n)
	var c = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	var am = make([]int, mult)
	var bm = make([]int, mult)
	var cm = make([]int, mult)

	for i := 0; i < n; i++ {
		am[a[i]%mult]++
		bm[b[i]%mult]++
		cm[c[i]%mult]++
	}

	sum := 0
	for i := 0; i < mult; i++ {
		for j := 0; j < mult; j++ {
			for k := 0; k < mult; k++ {
				if (i+j+k)%mult == 0 {
					sum += am[i] * bm[j] * cm[k]
				}
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", sum)
}
