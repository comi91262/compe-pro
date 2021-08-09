package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	defer writer.Flush()
	var w int
	fmt.Fscan(reader, &w)
	var h int
	fmt.Fscan(reader, &h)

	g := gcd(w, h)

	if w/g == 4 && h/g == 3 {
		fmt.Fprintf(writer, "%v\n", "4:3")
	} else if w/g == 16 && h/g == 9 {
		fmt.Fprintf(writer, "%v\n", "16:9")
	}
}
