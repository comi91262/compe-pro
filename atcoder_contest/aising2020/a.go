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

	var l int
	fmt.Fscan(reader, &l)
	var r int
	fmt.Fscan(reader, &r)
	var d int
	fmt.Fscan(reader, &d)
	ans := 0

	for i := l; i <= r; i++ {
		if i%d == 0 {
			ans++
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
