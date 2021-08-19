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

	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	ans := 0
	switch x {
	case 3:
		ans += 100000
	case 2:
		ans += 200000
	case 1:
		ans += 300000
	}

	switch y {
	case 3:
		ans += 100000
	case 2:
		ans += 200000
	case 1:
		ans += 300000
	}

	if x == 1 && y == 1 {
		ans += 400000
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
