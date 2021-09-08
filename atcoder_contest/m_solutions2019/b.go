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

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	m := map[string]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	if m["x"] > 7 {
		fmt.Fprintf(writer, "%v\n", "NO")
	} else {
		fmt.Fprintf(writer, "%v\n", "YES")
	}
}
