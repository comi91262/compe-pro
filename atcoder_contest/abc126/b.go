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

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	a, _ := strconv.Atoi(s[0] + s[1])
	b, _ := strconv.Atoi(s[2] + s[3])

	if 1 <= a && a <= 12 && 1 <= b && b <= 12 {
		fmt.Fprintf(writer, "%v\n", "AMBIGUOUS")
	} else if 1 <= a && a <= 12 {
		fmt.Fprintf(writer, "%v\n", "MMYY")
	} else if 1 <= b && b <= 12 {
		fmt.Fprintf(writer, "%v\n", "YYMM")
	} else {
		fmt.Fprintf(writer, "%v\n", "NA")
	}

}
