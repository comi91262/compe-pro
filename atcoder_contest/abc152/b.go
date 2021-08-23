package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	ans := ""

	if a >= b {
		for i := 0; i < a; i++ {
			ans += strconv.Itoa(b)
		}
		fmt.Fprintf(writer, "%v\n", ans)
		return
	}

	if a < b {
		for i := 0; i < b; i++ {
			ans += strconv.Itoa(a)
		}
		fmt.Fprintf(writer, "%v\n", ans)
		return
	}

}
