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

	var ss string
	fmt.Fscan(reader, &ss)
	var tt string
	fmt.Fscan(reader, &tt)

	s := strings.Split(ss, "")
	t := strings.Split(tt, "")

	ok := false
	for i := len(s) - 1; i >= 0; i-- {
		flag := true
		for j := 0; j < len(t); j++ {
			if i+j >= len(s) {
				flag = false
				break
			}
			if s[i+j] != t[j] && s[i+j] != "?" {
				flag = false
				break
			}
		}
		if flag {
			for j := 0; j < len(t); j++ {
				s[i+j] = t[j]
				ok = true
			}
		}

		if ok {
			break
		}
	}

	if !ok {
		fmt.Fprintf(writer, "%v\n", "UNRESTORABLE")
		return
	}

	for i := range s {
		if s[i] == "?" {
			fmt.Fprintf(writer, "%v", "a")
			continue
		}
		fmt.Fprintf(writer, "%v", s[i])
	}
	fmt.Fprintf(writer, "\n")
}
