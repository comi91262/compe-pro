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

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "YES")

}
