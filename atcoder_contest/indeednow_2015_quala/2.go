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
	var n int
	fmt.Fscan(reader, &n)

	var s = "indeednow"
	var ss = strings.Split(s, "")
	sort.Strings(ss)
	s = strings.Join(ss, "")

	for i := 0; i < n; i++ {
		var tt string
		fmt.Fscan(reader, &tt)
		var t = strings.Split(tt, "")
		sort.Strings(t)
		if s == strings.Join(t, "") {
			fmt.Fprintf(writer, "%v\n", "YES")
		} else {
			fmt.Fprintf(writer, "%v\n", "NO")
		}
	}
}
