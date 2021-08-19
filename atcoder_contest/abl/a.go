package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var k int
	fmt.Fscan(reader, &k)

	ans := ""
	for i := 0; i < k; i++ {
		ans += "ACL"
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
