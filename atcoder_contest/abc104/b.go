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

	if s[0] != "A" {
		fmt.Fprintf(writer, "%v\n", "WA")
		return
	}

	m := map[string]int{}
	for i := 2; i < len(s)-1; i++ {
		m[s[i]]++
	}

	if m["C"] != 1 {
		fmt.Fprintf(writer, "%v\n", "WA")
		return
	}

	for i := 0; i < len(s); i++ {
		if i == 0 {
			continue
		}
		if 2 <= i && i < len(s)-1 && s[i] == "C" {
			continue
		}
		if "A" <= s[i] && s[i] <= "Z" {
			fmt.Fprintf(writer, "%v\n", "WA")
			return

		}

	}
	fmt.Fprintf(writer, "%v\n", "AC")

}
