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
	var op string
	fmt.Fscan(reader, &op)
	var b int
	fmt.Fscan(reader, &b)

	switch op {
	case "+":
		fmt.Fprintf(writer, "%v\n", a+b)
	case "-":
		fmt.Fprintf(writer, "%v\n", a-b)
	case "*":
		fmt.Fprintf(writer, "%v\n", a*b)
	case "/":
		if b == 0 {
			fmt.Fprintf(writer, "%v\n", "error")
			return
		}
		fmt.Fprintf(writer, "%v\n", a/b)
	case "?":
		fmt.Fprintf(writer, "%v\n", "error")
	case "=":
		fmt.Fprintf(writer, "%v\n", "error")
	case "!":
		fmt.Fprintf(writer, "%v\n", "error")
	}
}
