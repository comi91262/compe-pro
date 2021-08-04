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

	var s string
	fmt.Fscan(reader, &s)

	for i := 0; i < a; i++ {
		if 48 > s[i] || s[i] > 57 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	if s[a] != "-"[0] {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	for i := a + 1; i < a+b+1; i++ {
		if 48 > s[i] || s[i] > 57 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "Yes")

}
