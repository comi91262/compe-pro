package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := ""
	for {
		if strings.HasSuffix(s, "dreamer"+t) {
			t = "dreamer" + t

		} else if strings.HasSuffix(s, "dream"+t) {
			t = "dream" + t

		} else if strings.HasSuffix(s, "eraser"+t) {
			t = "eraser" + t

		} else if strings.HasSuffix(s, "erase"+t) {
			t = "erase" + t
		} else {
			if s == t {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
			break
		}
	}

}package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var c string
	fmt.Fscan(reader, &c)

	switch c {
	case "a":
		fmt.Fprintf(writer, "%v\n", "vowel")
	case "i":
		fmt.Fprintf(writer, "%v\n", "vowel")
	case "u":
		fmt.Fprintf(writer, "%v\n", "vowel")
	case "e":
		fmt.Fprintf(writer, "%v\n", "vowel")
	case "o":
		fmt.Fprintf(writer, "%v\n", "vowel")
	default:
		fmt.Fprintf(writer, "%v\n", "consonant")
	}
}
