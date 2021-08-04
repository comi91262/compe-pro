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
	var s = strings.Split(ss, "")
	var t = strings.Split("WBWBWWBWBWBW", "")

	idx := 0
	for i := 0; i < len(s); i++ {
		ok := true
		for j := 0; j < len(t); j++ {
			if len(s) <= i+j {
				break
			}
			if s[i+j] != t[j] {
				ok = false
				break
			}
		}
		if ok {
			idx = i
			break
		}
	}

	u := []string{"Do", "Si", "", "La", "", "So", "", "Fa", "Mi", "", "Re", "", "Do"}

	fmt.Fprintf(writer, "%v\n", u[idx%12])
}
