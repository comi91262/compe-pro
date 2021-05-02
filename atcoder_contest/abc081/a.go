package main

import "fmt"

// func range(input string){
//     for _, c := range input {
//         fmt.Printf("%c ", c)
//         fmt.Println(reflect.TypeOf(c))
//     }
//     fmt.Println
// }

func main() {
	var s string
	fmt.Scan(&s)

	n := 0
	for _, c := range s {
		if c == '1' {
			n++
		}
	}

	fmt.Println(n)
}