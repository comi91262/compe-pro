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

	var a = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		a[i], _ = strconv.Atoi(s[i])
	}

	ans := 0
	for i := 0; i < 1<<(len(s)-1); i++ {
		sum := a[0]
		for j := 0; j < len(s)-1; j++ {
			if i&(1<<j) != 0 {
				ans += sum
				sum = 0
			}

			sum *= 10
			sum += a[j+1]
		}
		ans += sum
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
