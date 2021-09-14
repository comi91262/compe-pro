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
	r := regexp.MustCompile(`\d+`)
	fmt.Fprintf(writer, "%v\n", r.FindAllStringSubmatch(s, -1)[0][0])
}
