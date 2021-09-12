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

	ans := "1"
	for i := 0; i < n-1; i++ {
		ans += "0"
	}
	fmt.Fprintf(writer, "%v7\n", ans)

}
