package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value       string
	left, right *Node
}

func newNode(x string) *Node {
	p := new(Node)
	p.value = x
	return p
}

func searchNode(node *Node, x string) bool {
	for node != nil {
		switch {
		case x == node.value:
			return true
		case x < node.value:
			node = node.left
		case x > node.value:
			node = node.right
		default:
			return false
		}
	}
	return false
}

func insertNode(node *Node, x string) *Node {
	switch {
	case node == nil:
		return newNode(x)
	case x == node.value:
		return node
	case x < node.value:
		node.left = insertNode(node.left, x)
	case x > node.value:
		node.right = insertNode(node.right, x)
	default:
		return nil
	}
	return node
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}

	var t *Node
	for i := 0; i < n; i++ {
		if !searchNode(t, s[i]) {
			t = insertNode(t, s[i])
			fmt.Fprintf(writer, "%d\n", i+1)
		}
	}
}
