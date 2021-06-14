package main

var deque [20000000]int
var cl, cr = 10000000, 10000000

func addFront() {
	cl--
	deque[cl] = x[i]
}

func addBack() {
	deque[cr] = x[i]
	cr++
}

func removeFront() {
	cl++
}

func removeBack() {
	cr--
}

func get(x int) int {
	return deque[cl+x-1]
}
