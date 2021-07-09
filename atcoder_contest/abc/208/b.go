package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const f10 = 3628800

func main() {
	defer writer.Flush()

	var p int
	fmt.Fscan(reader, &p)

	ans := 0
	base := f10

	for i := 10; i > 0; i-- {
		ans += p / base
		p %= base
		base /= i
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
