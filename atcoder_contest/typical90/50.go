package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 1_000_000_000 + 7

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, l int
	fmt.Fscan(reader, &n, &l)

	var a = make([]int, n+1)

	a[0] = 1
	for i := 0; i < n+1; i++ {
		if i-1 >= 0 {
			a[i] += a[i-1]
		}
		if i-l >= 0 {
			a[i] += a[i-l]
		}

		a[i] = a[i] % MAX
	}

	fmt.Fprintf(writer, "%v\n", a[n])

}
