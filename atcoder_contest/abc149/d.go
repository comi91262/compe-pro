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
	var k int
	fmt.Fscan(reader, &k)
	var r int
	fmt.Fscan(reader, &r)
	var s int
	fmt.Fscan(reader, &s)
	var p int
	fmt.Fscan(reader, &p)

	var ss string
	fmt.Fscan(reader, &ss)
	var t = strings.Split(ss, "")

	ans := 0

	a := make([]string, n)
	for i := 0; i < n; i++ {
		switch t[i] {
		case "r":
			if i-k < 0 || i-k >= 0 && a[i-k] != "p" {
				ans += p
				a[i] = "p"
				continue
			}

			if i-k >= 0 && t[i-k] == "p" {
				if i+k < n {
					switch t[i-k] {
					case "r":
						a[i] = "s"
					case "s":
						a[i] = "r"
					case "p":
						a[i] = "r"
					}
				} else {
					a[i] = "r"
				}
			}
		case "s":
			if i-k < 0 || i-k >= 0 && a[i-k] != "r" {
				ans += r
				a[i] = "r"
				continue
			}

			if i-k >= 0 && t[i-k] == "r" {
				if i+k < n {
					switch t[i-k] {
					case "p":
						a[i] = "s"
					case "s":
						a[i] = "p"
					case "r":
					}
				} else {
					a[i] = "r"
				}
			}
		case "p":
			if i-k < 0 || i-k >= 0 && a[i-k] != "s" {
				ans += s
				a[i] = "s"
				continue
			}

			if i-k >= 0 && t[i-k] == "s" {
				if i+k < n {
					switch t[i-k] {
					case "r":
						a[i] = "p"
					case "p":
						a[i] = "r"
					case "s":
					}
				} else {
					a[i] = "r"
				}
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
