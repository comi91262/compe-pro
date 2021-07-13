package main

import (
	"bufio"
	"fmt"
	"os"
)

var a [100][6]int
var b [101]int

const MAX = 1_000_000_000 + 7

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	b[0] = 1
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < 6; j++ {
			sum += a[i][j]
		}
		b[i+1] = b[i] * sum % MAX
	}

	fmt.Fprintf(writer, "%v\n", b[n])

}
