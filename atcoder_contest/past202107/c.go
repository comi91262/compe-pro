package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n string
	fmt.Fscan(reader, &n)
	var l int
	fmt.Fscan(reader, &l)
	var r int
	fmt.Fscan(reader, &r)

	m, _ := strconv.Atoi(n)
	if n == "0" {
		if l <= m && m <= r {
			fmt.Fprintf(writer, "%v\n", "Yes")
		} else {
			fmt.Fprintf(writer, "%v\n", "No")
		}
		return
	}

	if strings.HasPrefix(n, "0") {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	if l <= m && m <= r {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
