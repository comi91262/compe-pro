package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a string
	fmt.Fscan(reader, &a)
	var b string
	fmt.Fscan(reader, &b)
	var c string
	fmt.Fscan(reader, &c)

	var ss string
	fmt.Fscan(reader, &ss)
	var t = strings.Split(ss, "")
	for i := 0; i < len(t); i++ {
		switch t[i] {
		case "1":
			fmt.Fprintf(writer, "%v", a)
		case "2":
			fmt.Fprintf(writer, "%v", b)
		case "3":
			fmt.Fprintf(writer, "%v", c)
		}
	}
	fmt.Fprintf(writer, "\n")

}
