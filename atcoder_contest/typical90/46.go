package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	var a2 = make([]int, 46)
	var b = make([]int, n)
	var b2 = make([]int, 46)
	var c = make([]int, n)
	var c2 = make([]int, 46)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i] = a[i] % 46
		a2[a[i]]++
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
		b[i] = b[i] % 46
		b2[b[i]]++
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
		c[i] = c[i] % 46
		c2[c[i]]++
	}

	sum := 0
	for i := 0; i < len(a2); i++ {
		for j := 0; j < len(b2); j++ {
			for k := 0; k < len(c2); k++ {
				if a2[i] > 0 && b2[j] > 0 && c2[k] > 0 {
					if (i+j+k)%46 == 0 {
						sum += a2[i] * b2[j] * c2[k]
					}
				}
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", sum)
}
