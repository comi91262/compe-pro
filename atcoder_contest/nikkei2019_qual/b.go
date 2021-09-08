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
	var ss string
	fmt.Fscan(reader, &ss)
	var a = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var b = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var c = strings.Split(ss, "")

	cnt := 0
	for i := 0; i < n; i++ {
		m := map[string]int{}
		m[a[i]]++
		m[b[i]]++
		m[c[i]]++
		cnt += len(m) - 1
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
