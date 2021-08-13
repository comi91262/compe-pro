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

	var t = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		t[i], _ = strconv.Atoi(s[i])
	}

	ans := 0
	pre := t[0]
	for i := 1; i < len(s); i++ {
		if pre == t[i] {
			ans++
			pre = 1 - t[i]
			continue
		}

		pre = t[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
