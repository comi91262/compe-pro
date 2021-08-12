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
	var x = make([]float64, n)
	var u = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &u[i])
	}
	ans := 0.0
	for i := 0; i < n; i++ {
		switch u[i] {
		case "JPY":
			ans += x[i]
		case "BTC":
			ans += x[i] * 380000.0
		}

	}
	fmt.Fprintf(writer, "%v\n", ans)

}
