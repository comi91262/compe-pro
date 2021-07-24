package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var tmp string
	fmt.Fscan(reader, &tmp)
	s := strings.Split(tmp, "")
	sort.Strings(s)

	if len(s) == 1 && s[0] == "a" {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}
	fmt.Fprintf(writer, "%v\n", "a")
}
