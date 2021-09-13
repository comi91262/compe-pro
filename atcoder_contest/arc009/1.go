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
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		var b int
		fmt.Fscan(reader, &b)
		ans += a * b
	}
	fmt.Fprintf(writer, "%v\n", ans*105/100)
}
