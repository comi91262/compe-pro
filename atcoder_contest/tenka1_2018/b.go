package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var k int
	fmt.Fscan(reader, &k)
	for i := 0; i < k; i++ {
		if i%2 == 0 {
			if a%2 != 0 {
				a--
			}
			b += a / 2
			a /= 2
		} else {
			if b%2 != 0 {
				b--
			}
			a += b / 2
			b /= 2
		}
	}
	fmt.Fprintf(writer, "%v %v\n", a, b)
}
