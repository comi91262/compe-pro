package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	m := map[string]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	cnt := 1
	for _, v := range m {
		cnt *= (v + 1)
		cnt %= mod
	}

	fmt.Fprintf(writer, "%v\n", (cnt-1)%mod)
}
