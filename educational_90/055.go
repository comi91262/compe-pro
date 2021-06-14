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

	var n, p, q int

	fmt.Fscan(reader, &n, &p, &q)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					for m := l + 1; m < n; m++ {
						if a[i]*a[j]%p*a[k]%p*a[l]%p*a[m]%p == q {
							sum++
						}
					}
				}
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", sum)

}
