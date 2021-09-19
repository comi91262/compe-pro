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

	i := ""
	mx := 0
	for k, v := range m {
		if mx < v {
			mx = v
			i = k
		}
	}
	fmt.Fprintf(writer, "%v\n", i)

}
