package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func dfs(n, cur, used int) int {
	if cur > n {
		return 0
	}

	r := 0
	if used == 0b111 {
		r++
	}

	r += dfs(n, cur*10+7, used|0b001)
	r += dfs(n, cur*10+5, used|0b010)
	r += dfs(n, cur*10+3, used|0b100)

	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	fmt.Fprintf(writer, "%v\n", dfs(n, 0, 0))
}
