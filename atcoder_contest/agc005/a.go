package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type stack []string

func (s *stack) push(n string) {
	*s = append(*s, n)
}

func (s *stack) pop() string {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func (s *stack) front() string {
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	ans := 0
	st := stack{}
	for i := 0; i < len(s); i++ {
		if s[i] == "S" {
			st.push(s[i])
		} else {
			if !st.empty() {
				st.pop()
				ans++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", len(s)-ans*2)
}
