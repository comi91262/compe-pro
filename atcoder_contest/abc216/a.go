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
	a := strings.Split(s, ".")
	y, _ := strconv.Atoi(a[1])

	switch {
	case 0 <= y && y <= 2:
		fmt.Fprintf(writer, "%v\n", a[0]+"-")
	case 3 <= y && y <= 6:
		fmt.Fprintf(writer, "%v\n", a[0])
	case 7 <= y && y <= 9:
		fmt.Fprintf(writer, "%v\n", a[0]+"+")
	}

}
