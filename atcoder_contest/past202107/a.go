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

	a := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		a[i], _ = strconv.Atoi(s[i])
	}

	o := 0
	e := 0
	c := 0
	for i := 0; i < len(a); i++ {
		if i == len(a)-1 {
			c = a[i]
			continue
		}
		if i%2 == 0 {
			o += a[i]
		} else {
			e += a[i]
		}
	}

	if (o*3+e)%10 == c {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
