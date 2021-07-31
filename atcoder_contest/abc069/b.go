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

	var s string
	fmt.Fscan(reader, &s)

	t := strings.Split(s, "")

	fmt.Fprintf(writer, "%v\n", t[0]+strconv.Itoa(len(t)-2)+t[len(t)-1])

}
