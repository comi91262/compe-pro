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

	var n, a, b, c int
	fmt.Fscan(reader, &n, &a, &b, &c)

	const max = 9999
	min := 10000
	for i := 0; i <= max; i++ {
		for j := 0; j+i <= max; j++ {
			var rest = n - a*i - b*j
			var k = rest / c
			if rest >= 0 && rest%c == 0 && min > i+j+k {
				min = i + j + k
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", min)
}
