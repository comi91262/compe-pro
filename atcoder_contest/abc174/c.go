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

	var k int
	fmt.Fscan(reader, &k)

	ans := 1
	n := 7 % k
	for {
		if n == 0 {
			break
		}
		n = (n*10 + 7) % k

		ans++
		if ans > k {
			ans = -1
			break
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
