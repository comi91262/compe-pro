package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)
	r := regexp.MustCompile(`C[A-Z]*F`)

	if r.MatchString(s) {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
