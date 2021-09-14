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

	ans := 0
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(s[i])
		if i%2 == 0 {
			ans += n
		} else {
			ans -= n
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
