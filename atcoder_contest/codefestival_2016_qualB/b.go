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
	var n int
	fmt.Fscan(reader, &n)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	abn := 0
	bn := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "a":
			if abn < a+b {
				fmt.Fprintf(writer, "%v\n", "Yes")
				abn++
			} else {
				fmt.Fprintf(writer, "%v\n", "No")
			}
		case "b":
			if abn < a+b && bn < b {
				fmt.Fprintf(writer, "%v\n", "Yes")
				abn++
				bn++
			} else {
				fmt.Fprintf(writer, "%v\n", "No")
			}
		case "c":
			fmt.Fprintf(writer, "%v\n", "No")
		}
	}

}

