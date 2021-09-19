package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()
	var s string
	fmt.Fscan(reader, &s)
	var t string
	fmt.Fscan(reader, &t)

	isCase := false
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] && abs(int(s[i])-int(t[i])) != 32 {
			fmt.Fprintf(writer, "%v\n", "different")
			return
		}
		if abs(int(s[i]-t[i])) == 32 {
			isCase = true
		}
	}
	if isCase {
		fmt.Fprintf(writer, "%v\n", "case-insensitive")
	} else {
		fmt.Fprintf(writer, "%v\n", "same")
	}

}
