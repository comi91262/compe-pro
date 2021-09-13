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
	r := regexp.MustCompile(`(i|I)[A-Za-z]*(c|C)[A-Za-z]*(t|T)`)

	var s string
	fmt.Fscan(reader, &s)

	if r.MatchString(s) {
		fmt.Fprintf(writer, "%v\n", "YES")
	} else {
		fmt.Fprintf(writer, "%v\n", "NO")
	}

}
