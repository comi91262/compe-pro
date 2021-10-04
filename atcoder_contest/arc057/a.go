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
	var k int
	fmt.Fscan(reader, &k)

	if k == 0 {
		fmt.Fprintf(writer, "%v\n", 2_000_000_000_000-a)
		return
	}

	day := 0
	tot := a
	for tot+1+tot*k < 2_000_000_000_000 {
		day++
		tot += 1 + tot*k
	}
	fmt.Fprintf(writer, "%v\n", day+1)
}
