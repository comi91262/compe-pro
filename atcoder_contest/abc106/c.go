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
	var k int
	fmt.Fscan(reader, &k)

	s := strings.Split(ss, "")

	cnt := 0
	for i := range s {
		if cnt == k {
			break
		}
		if s[i] == "1" {
			cnt++
			continue
		}

		fmt.Fprintf(writer, "%v\n", s[i])
		return
	}

	fmt.Fprintf(writer, "%v\n", "1")
}
