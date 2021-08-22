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

	if b == 1 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}
	ans := 1
	cur := a
	for cur < b {
		cur += (a - 1)
		ans++
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
