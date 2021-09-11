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

	var p = make([]int, 26)
	for i := 0; i < 26; i++ {
		fmt.Fscan(reader, &p[i])
	}

	ans := make([]string, 26)
	for i := 0; i < 26; i++ {
		ans[i] = string("a"[0] + byte(p[i]-1))
	}
	fmt.Fprintf(writer, "%v\n", strings.Join(ans, ""))
}

