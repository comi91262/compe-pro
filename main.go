package main

import (
	"bufio"
	"os"
)

const mod = 1_000_000_000 + 7

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

}
