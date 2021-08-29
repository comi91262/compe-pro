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

	ans := []string{}
	for n > 0 {
		if n%2 == 1 {
			ans = append(ans, "A")
			n--
		} else {
			ans = append(ans, "B")
			n /= 2
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}

	fmt.Fprintf(writer, "%v\n", strings.Join(ans, ""))

}
