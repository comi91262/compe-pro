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

	ans := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			continue
		}
		if i%5 == 0 {
			continue
		}

		ans += i
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
