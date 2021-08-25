package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	s := map[string]int{}
	t := map[string]int{}
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)

		if strings.HasPrefix(tmp, "!") {
			s[tmp]++
		} else {
			t[tmp]++
		}
	}

	for k := range t {
		if s["!"+k] > 0 {
			fmt.Fprintf(writer, "%v\n", k)
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "satisfiable")
}
