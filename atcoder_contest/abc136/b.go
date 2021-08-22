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

	var n int
	fmt.Fscan(reader, &n)

	ans := 0
	for i := 1; i <= n; i++ {
		t := i
		d := 0
		for 0 < t {
			d++
			t /= 10
		}
		if d%2 != 0 {
			ans++
		}

	}
	fmt.Fprintf(writer, "%v\n", ans)
}
