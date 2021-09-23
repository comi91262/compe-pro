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

	if n == 1 {
		fmt.Fprintf(writer, "%v\n", "Not Prime")
		return
	}

	if n%10%2 != 0 && n%10%5 != 0 && n%3 != 0 {
		fmt.Fprintf(writer, "%v\n", "Prime")
		return
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Fprintf(writer, "%v\n", "Not Prime")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "Prime")

}
