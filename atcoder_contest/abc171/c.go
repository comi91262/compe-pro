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

	d := []int{}
	for n > 0 {
		n--
		d = append(d, n%26)
		n /= 26
	}

	for i := 0; i < len(d)/2; i++ {
		d[i], d[len(d)-i-1] = d[len(d)-i-1], d[i]
	}

	ans := make([]string, len(d))
	for i := 0; i < len(d); i++ {
		ans[i] = string(byte(d[i]) + "a"[0])
	}

	fmt.Fprintf(writer, "%v\n", strings.Join(ans, ""))

}
